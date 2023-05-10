package GoDataExtractor

import (
	log "github.com/sirupsen/logrus"
	"io"
	"reflect"
)

// Do func extract the data from file and put it into the ConstructedOutput chan
// use go Do() to make it as a new runtime.
func (p *DataExtractor) Do() {
	Scanner := p.scanner
	Selectors := p.docQueries

	Scanner.InitIndex(Selectors)
	defer p.scanner.Close()
	defer close(p.ConstructedOutput)

	for {
		var recordBuilder = p.targetStruct
		recordBuilderT := reflect.TypeOf(p.targetStruct).Elem()
		recordBuilderV := reflect.ValueOf(p.targetStruct).Elem()

		if p.lastErr != nil {
			goto FallBack
		}

		// for all selectors to build recordBuilder
		p.lastErr = Scanner.Select(Selectors)
		// if read EOF or other error in this Loop
		if p.lastErr != nil {
			goto FallBack
		}

		for _, v := range Selectors {
			for i := 0; i < recordBuilderT.NumField(); i++ {
				CurrentTagStr := recordBuilderT.Field(i).Tag.Get("extract")
				if CurrentTagStr == v.TagName {
					// copy Value to dest
					recordBuilderV.Field(i).Set(reflect.ValueOf(v.Value))
				}
			}
			//reflect.TypeOf(recordBuilder).
			//	reflect.ValueOf(recordBuilder).Elem().FieldByName(v.TagName).Set(reflect.ValueOf(v.GetValue()))
			// recordBuilder = // Todo: Build T type from Selected values with TagName
		}

		// Ask for First and Second Record
		if p.count == 0 || p.count == 1 {
			log.Infof("This is No.%v Record generated struct.", p.count)
		}

		p.targetStruct = recordBuilder
		p.lastErr = p.customFunc(p)
		if p.lastErr != nil {
			log.Error("CustomFunc Error: ", p.lastErr)
			goto FallBack
		}

		p.send()

		if p.count%100 == 0 {
			log.Info("Record Count: ", p.count)
		}
		p.count++
		continue

		// ToDo: Get Another parsing Error and Record them and continue the parsing (non break) This break will let everything die.
	FallBack:
		log.Debug("Receive Error", p.lastErr)
		if p.lastErr == io.EOF {
			log.Info("End Record At No.", p.count, " line")
			log.Info("Got End Of file")
			break // stop parsing and exit.
		} // make fallback
		p.lastErr = p.fallbackFunc(p)
	}
}
