package db

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"

	"app/GoSample/infra/constant"
	"app/GoSample/logger"
)

//Additional init script that contains CRUD operations, SP's etc.
func InitScripts(db *sql.DB) {
	initScriptPath := os.Getenv("InitSQLFilePath")
	initScriptFile, err := ioutil.ReadFile(initScriptPath)
	if err != nil {
		logger.ErrorLog("An error occured while reading init.sql file - dbInitializer.go - Error:", err.Error())
	}
	initScript := string(initScriptFile)

	transaction, err := db.Begin()
	if err != nil {
		logger.ErrorLog("An error occured while beginning transaction - dbInitializer.go - Error:", err.Error())
	} else {
		logger.TransactionLog("Transaction began")
	}

	defer func() {
		err := transaction.Rollback()
		if err != nil {
			logger.ErrorLog("An error occured while rollbacking transaction - dbInitializer.go - Error:", err.Error())
		} else {
			logger.TransactionLog("Transaction rollback")
		}
	}()

	for _, statement := range strings.Split(initScript, constant.NextLine) {
		statement := strings.TrimSpace(statement)
		if statement == constant.EmptyString {
			continue
		}
		if _, err := transaction.Exec(statement); err != nil {
			logger.ErrorLog("An error occured while executing statements - dbInitializer.go - Error:", err.Error())
		} else {
			logger.TransactionLog(statement)
		}
	}

	err = transaction.Commit()
	if err != nil {
		logger.ErrorLog("An error occured while committing transaction - dbInitializer.go - Error:", err.Error())
	} else {
		logger.TransactionLog("Transaction committed")
	}
}
