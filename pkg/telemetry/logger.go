package telemetry

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[OMNIPING] ", log.Ldate|log.Ltime|log.Lshortfile)

func SetLogLevel(debug bool) {
	if debug {
		Logger.Println("Debug logging enabled")
	}
}
