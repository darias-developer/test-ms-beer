package middleware

import (
	"log"
	"os"
)

var (
	LogWarn  *log.Logger
	LogInfo  *log.Logger
	LogError *log.Logger
)

func LoggerInit() {
	file, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	LogInfo = log.New(file, "[INFO ] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	LogWarn = log.New(file, "[WARN ] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
	LogError = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
}
