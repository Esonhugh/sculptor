package GoDataExtractor

import (
	log "github.com/sirupsen/logrus"
	"io"
)

func (p DataExtractor[T]) Do() {
	Scanner := p.scanner
	Selectors := p.docQueries

	Scanner.InitIndex(Selectors)
	defer p.scanner.Close()
	RecordCount := p.count
	for {
		var err = p.lastErr
		var recordBuilder T

		// for all selectors to build recordBuilder
		err = Scanner.Select(Selectors)
		// if read EOF or other error in this Loop
		if err != nil {
			p.lastErr = err
			log.Debug("Receive Select Error", err)
			log.Error(err.Error())
			if err == io.EOF {
				log.Info("End Record At No.", RecordCount, " line")
				log.Info("Got End Of file")
				break
			}
			// ToDo: Get Another parsing Error and Record them and continue the parsing (non break) This break will let everything die.
			break
		}

		/*
			for _,_ := range Selectors {
				recordBuilder = T{}// Todo: Build T type from Selected values with TagName
			}
		*/
		// Ask for First and Second Record
		if RecordCount == 0 || RecordCount == 1 {
			/*
				log.Infof("This is No.%v Record generated struct.", RecordCount)
				for _, info := range recordBuilder {
					parser.PrintBasicInfo(*info)
				}
				if !AskForSure("Is this Record Correct?") {
					log.Panicf("User Deny to send the Record %d", RecordCount)
				}
			*/
		}

		p.ConstructedOutput <- recordBuilder

		if RecordCount%100 == 0 {
			log.Info("Record Count: ", RecordCount)
		}
		RecordCount++
	}
}
