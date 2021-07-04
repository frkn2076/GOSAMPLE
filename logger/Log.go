package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

func Start() {
	info = initFileLogger("InfoLog")
	error = initFileLogger("ErrorLog")
	transaction = initFileLogger("TransactionLog")
	QueryLogger = initQueryLogger("TransactionLog")
}

var info *log.Logger
var error *log.Logger
var transaction *log.Logger
var QueryLogger logger.Interface

func ErrorLog(logText ...interface{}) {
	error.Println(logText)
}

func InfoLog(logText ...interface{}) {
	info.Println(logText)
}

func TransactionLog(logText ...interface{}) {
	transaction.Println(logText)
}

func initFileLogger(folderName string) *log.Logger {
	dt := time.Now()
	today := dt.Format("02-Jan-2006")
	loggerFilePath := os.Getenv("LoggerFilePath")
	fileName := fmt.Sprintf(loggerFilePath, folderName, today)

	//check log file created before
	_, err := os.Stat(fileName)
	fileNotExist := os.IsNotExist(err)

	if fileNotExist {
		folderPath := strings.TrimSuffix(loggerFilePath, "/%s.log")
		formattedFolderPath := fmt.Sprintf(folderPath, folderName)
		os.MkdirAll(formattedFolderPath, 0700) // Create your folder
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "prefix: ", log.LstdFlags)
	if fileNotExist {
		logger.Println(folderName, "has created")
	}
	return logger
}

func initQueryLogger(folderName string) logger.Interface {
	transactionLogger := initFileLogger(folderName)
	newLogger := logger.New(
		transactionLogger,
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)
	return newLogger
}
