package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	// "gorm.io/gorm/logger"
)

var info *log.Logger = initFileLogger("InfoLog")
var error *log.Logger = initFileLogger("ErrorLog")
var service *log.Logger = initFileLogger("ServiceLog")
var transaction *log.Logger = initFileLogger("TransactionLog")
// var QueryLogger logger.Interface = initQueryLogger("TransactionLog")

func ErrorLog(logText ...interface{}) {
	error.Println(logText)
}

func InfoLog(logText ...interface{}) {
	info.Println(logText)
}

func ServiceLog(logText ...interface{}) {
	service.Println(logText)
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

// removed since gorm unsupport versions
// func initQueryLogger(folderName string) logger.Interface {
// 	transactionLogger := initFileLogger(folderName)
// 	newLogger := logger.New(
// 		transactionLogger,
// 		logger.Config{
// 			SlowThreshold:             time.Second, // Slow SQL threshold
// 			LogLevel:                  logger.Info, // Log level
// 			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
// 			Colorful:                  false,       // Disable color
// 		},
// 	)
// 	return newLogger
// }

