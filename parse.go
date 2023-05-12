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

	d.debugInfo()

	Scanner.InitIndex(Selectors)

	defer d.scanner.Close()
	defer close(d.ConstructedOutput)

	for {
		// var recordBuilder = d.targetStruct
		recordBuilderT := reflect.TypeOf(d.targetStruct).Elem()
		recordBuilderV := reflect.ValueOf(d.targetStruct).Elem()

		if d.Error() != nil {
			goto FallBack
		}

		// for all selectors to build recordBuilder
		d.lastErr = Scanner.Select(Selectors)
		// if read EOF or other error in this Loop
		if d.Error() != nil {
			log.WithFields(log.Fields{
				"Count":                d.count,
				"err":                  d.Error(),
				"during SelectingLine": d.scanner.CurrentLine(),
			}).Trace("Scanner selecting porcess fails.")
			goto FallBack
		}

		for _, v := range Selectors {
			for i := 0; i < recordBuilderT.NumField(); i++ {
				CurrentTagStr := recordBuilderT.Field(i).Tag.Get(d.options.TagKey)
				log.WithFields(log.Fields{
					"Count":              d.count,
					"SelectingStructTag": CurrentTagStr,
					"SelectorTag":        v.TagName,
				}).Trace("Current Select tag match status")
				if CurrentTagStr == v.TagName {
					// copy Value to dest
					log.WithFields(log.Fields{
						"Count":              d.count,
						"SelectingStructTag": CurrentTagStr,
						"SelectorTag":        v.TagName,
					}).Trace("current tags is matched. Setting value.")
					recordBuilderV.Field(i).Set(reflect.ValueOf(v.Value))
				}
			}
			// Todo: Build T type from Selected values with TagName
		}

		// Ask for First and Second Record
		if d.count == 0 || d.count == 1 {
			log.WithFields(log.Fields{
				"Count":  d.count,
				"Record": d.targetStruct,
			}).Info("Record generated struct. ")
		} else {
			log.WithFields(log.Fields{
				"Count":  d.count,
				"Record": d.targetStruct,
			}).Trace("Record generated struct. ")
		}

		// d.targetStruct = recordBuilder
		for i, f := range d.customFunc {
			d.lastErr = f(d)
			if d.Error() != nil {
				log.WithFields(log.Fields{
					"Count":           d.count,
					"CustomFuncIndex": i,
					"err":             d.Error(),
				}).Errorf("CustomFunc Exec with Error")
				goto FallBack
			}
		}

		log.WithFields(log.Fields{
			"Count":  d.count,
			"Record": d.targetStruct,
		}).Trace("Preprocess with customFunc successfully. Sending Data")

		d.send()

		if d.count%100 == 0 && d.count != 0 {
			log.Infof("%v Records have been generated.", d.count)
		}
		if d.count%1000 == 0 && d.count > 1000 {
			log.Infof("%v Records have been generated.", d.count)
		}

		d.count++
		continue

	FallBack:
		// ToDo: Get Another parsing Error and Record them and continue the parsing (non break) This break will let everything die.
		log.WithFields(log.Fields{
			"Count": d.count,
			"err":   d.Error(),
		}).Debug("Receive Error and fallback init")

		if d.Error() == io.EOF {
			log.Info("End Record At No.", d.count, " line")
			log.Info("Got End Of file")
			break // stop parsing and exit.
		}
		// do fallbacks
		log.WithFields(log.Fields{
			"Count": d.count,
			"err":   d.Error(),
		}).Trace("Error is not caused by end of file.")
		for i, f := range d.fallbackFunc {
			err := f(d) // force fallbackFunc success.
			if err != nil {
				log.WithFields(log.Fields{
					"Count":                 d.count,
					"LastErr":               d.Error(),
					"FallbackFuncIndex":     i,
					"Error by FallbackFunc": err,
				}).Errorf("Fallback Error. Program will exit.")
				break
			}
		}
	}
	d.CTX.Done()
}
