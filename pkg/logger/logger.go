package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	logsDir = "logs"

	loggerInfo *log.Logger
	loggerError *log.Logger
)

func Init() error {

	var err error

	if err = os.MkdirAll(fmt.Sprintf("./%s", logsDir), 0766); err != nil {
		return err
	}


	if loggerInfo, err = newLogger("info"); err != nil {
		return err
	}

	if loggerError, err = newLogger("error"); err != nil {
		return err
	}

	return nil
}

func newLogger(fileName string) (*log.Logger, error) {
	file, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", fileName), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(file, "", log.Lshortfile|log.Ldate|log.Ltime)

	return logger, nil
}

func Info(f string, v ...interface{}) {
	//标准输出
	log.Printf(f, v ...)

	//日志记录
	loggerInfo.Printf(f, v ...)
}

func Error(f string, v ...interface{}) {
	//标准输出
	log.Printf(f, v ...)

	//日志记录
	loggerError.Printf(f, v ...)
}