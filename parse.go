package GoDataExtractor

import (
	log "github.com/sirupsen/logrus"
	"io"
	"reflect"
)

func (p DataExtractor[T]) Do() {
	Scanner := p.scanner
	Selectors := p.docQueries

	Scanner.InitIndex(Selectors)
	defer p.scanner.Close()

	for {
		var recordBuilder T

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
			reflect.ValueOf(recordBuilder).Elem().FieldByName(v.TagName).Set(reflect.ValueOf(v.GetValue()))
			// recordBuilder = // Todo: Build T type from Selected values with TagName
		}

		/*
			// Ask for First and Second Record
			if p.count == 0 || p.count == 1 {
					log.Infof("This is No.%v Record generated struct.", p.count)
					for _, info := range recordBuilder {
						parser.PrintBasicInfo(*info)
					}
					if !AskForSure("Is this Record Correct?") {
						log.Panicf("User Deny to send the Record %d", p.count)
					}
			}
		*/

		p.targetStruct = recordBuilder
		p.lastErr = p.CustomFunc()
		if p.lastErr != nil {
			log.Error("CustomFunc Error: ", p.lastErr)
			goto FallBack
		}

		p.ConstructedOutput <- recordBuilder

		if p.count%100 == 0 {
			log.Info("Record Count: ", p.count)
		}
		p.count++

		// ToDo: Get Another parsing Error and Record them and continue the parsing (non break) This break will let everything die.
	FallBack:
		log.Debug("Receive Select Error", p.lastErr)
		log.Error(p.lastErr.Error())
		if p.lastErr == io.EOF {
			log.Info("End Record At No.", p.count, " line")
			log.Info("Got End Of file")
			break
		}
		p.lastErr = p.FallbackFunc()
	}
}
