package decorator

import (
	"io"
	"log"
	"os"
)

var (
	Debug       *log.Logger
	Info        *log.Logger
	Error       *log.Logger
	InfoHandler io.Writer
)

func InitLog(fl string, dh, ih, eh io.Writer) {
	if len(fl) > 0 {
		_ = os.Remove(fl)
		file, err := os.OpenFile(fl, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to create %s log file : %v\n", fl, err)
		}

		dh = io.MultiWriter(file, dh)
		ih = io.MultiWriter(file, ih)
		eh = io.MultiWriter(file, eh)

		InfoHandler = ih

		Debug = log.New(dh, "DEBUG : ", log.Ldate|log.Ltime|log.Lshortfile)

		Info = log.New(ih, "INFO  : ",
			log.Ltime)

		Error = log.New(eh, "ERROR : ",
			log.Ldate|log.Ltime|log.Lshortfile)

	}
}
