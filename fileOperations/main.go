package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	FileLogPath = "logs"

	LogInfo("Selamlar")
}

var FileLogPath string

func LogInfo(message string) {

	logMessage(message, "INFO")
}

func LogError(message string) {

	logMessage(message, "ERROR")
}

func LogWarning(message string) {

	logMessage(message, "WARNING")
}

func logMessage(message string, logType string) {

	FileLogPath = fmt.Sprintf("%s%s%s", FileLogPath, "_", time.Now().Format("20060102"))

	message = fmt.Sprintf("- [%s] - %s", logType, message)

	f, err := os.OpenFile(FileLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)

}
