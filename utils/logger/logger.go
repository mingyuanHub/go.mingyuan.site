package logger

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func InitLogger() error {
	file, err := os.Create("./logs/app.log")
	if err != nil {
		return err
	}

	Logger = log.New(file, "", log.Llongfile|log.Ldate|log.Ltime)

	return nil
}

func Info(f string, v ...interface{}) {
	Logger.Printf(f, v)
}