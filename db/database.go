package db

import (
	"context"
	"database/sql"
	"os"
	"time"
	"io/ioutil"
	"strings"

	"app/GoSample/infra/constant"
	"app/GoSample/db/entities"
	"app/GoSample/logger"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgreDB *sql.DB
var MongoDB *mongo.Database
var GormDB *gorm.DB

func ConnectDatabases(){
	PostgreDB = initPostgreDB()
	MongoDB = initMongoDB()
	GormDB = initGorm()
}

func MigrateTables() {
	GormDB.AutoMigrate(&entities.Account{}, &entities.Localization{}, &entities.Todo{})
}

func InitScripts() {
	initScriptPath := os.Getenv("InitSQLFilePath")
	initScriptFile, err := ioutil.ReadFile(initScriptPath)
	if err != nil {
		logger.ErrorLog("An error occured while reading init.sql file - dbInitializer.go - Error:", err.Error())
	}
	initScript := string(initScriptFile)

	transaction, err := PostgreDB.Begin()
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

	logger.InfoLog("Init sql script has runned")
}

func initPostgreDB() *sql.DB {
	connection := os.Getenv("PGSQLConnection")
	db, err := sql.Open("postgres", connection)
	if err != nil {
		logger.ErrorLog("An error occured while postgre connection is establishing. - Error:", err.Error())
		os.Exit(0)
	}

	//Ping for 2 seconds
	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfunc()
	if err = db.PingContext(ctx); err != nil {
		logger.ErrorLog("An error occured while postgre ping:", err.Error())
	}
	logger.InfoLog("Postgre database connection is opened")
	return db
}

func initMongoDB() *mongo.Database {
	connection := os.Getenv("MongoConnection")
	clientOptions := options.Client().ApplyURI(connection)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logger.ErrorLog("An error occured while mongo connection is establishing ", err.Error())
		os.Exit(0)
	}

	//Ping for 2 seconds
	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfunc()

	if err = client.Connect(ctx); err != nil {
		logger.ErrorLog("An error occured while mongo ping:", err.Error())
	}
	logger.InfoLog("Mongo database connection is opened")

	logDB := client.Database("LogDB")
	return logDB
}

func initGorm() *gorm.DB {

	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: PostgreDB, // Initialize gorm with the existing db connection
		}),
		&gorm.Config{
			Logger:                 logger.QueryLogger,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		logger.ErrorLog("An error occured while gorm driver is establishing: ", err.Error())
		os.Exit(0)
	}

	return gormDB
}
