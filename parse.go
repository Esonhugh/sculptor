package sculptor

import (
	log "github.com/sirupsen/logrus"
	"io"
	"reflect"
)

// Do func extract the data from file and put it into the ConstructedOutput chan
// use go Do() to make it as a new runtime.
func (d *DataSculptor) Do() {
	Scanner := d.scanner
	Selectors := d.docQueries

	Scanner.InitIndex(Selectors)
	defer d.scanner.Close()
	defer close(d.ConstructedOutput)

	for {
		// var recordBuilder = d.targetStruct
		recordBuilderT := reflect.TypeOf(d.targetStruct).Elem()
		recordBuilderV := reflect.ValueOf(d.targetStruct).Elem()

		if d.lastErr != nil {
			goto FallBack
		}

		// for all selectors to build recordBuilder
		d.lastErr = Scanner.Select(Selectors)
		// if read EOF or other error in this Loop
		if d.lastErr != nil {
			goto FallBack
		}

		for _, v := range Selectors {
			for i := 0; i < recordBuilderT.NumField(); i++ {
				CurrentTagStr := recordBuilderT.Field(i).Tag.Get(d.options.TagKey)
				if CurrentTagStr == v.TagName {
					// copy Value to dest
					recordBuilderV.Field(i).Set(reflect.ValueOf(v.Value))
				}
			}
			// Todo: Build T type from Selected values with TagName
		}

		// Ask for First and Second Record
		if d.count == 0 || d.count == 1 {
			log.Infof("This is No.%v Record generated struct.", d.count)
		}

		// d.targetStruct = recordBuilder
		for i, f := range d.customFunc {
			d.lastErr = f(d)
			if d.lastErr != nil {
				log.Errorf("CustomFunc[%v]Error: ", i, d.lastErr)
				goto FallBack
			}
		}

		d.send()

		if d.count%100 == 0 {
			log.Info("Record Count: ", d.count)
		}
		d.count++
		continue

	FallBack:
		// ToDo: Get Another parsing Error and Record them and continue the parsing (non break) This break will let everything die.
		log.Debug("Receive Error", d.lastErr)
		if d.lastErr == io.EOF {
			log.Info("End Record At No.", d.count, " line")
			log.Info("Got End Of file")
			break // stop parsing and exit.
		}
		// do fallbacks
		for _, f := range d.fallbackFunc {
			err := f(d) // force fallbackFunc success.
			if err != nil {
				log.Error("Fallback Error: ", d.lastErr)
				break
			}
		}
	}
}
