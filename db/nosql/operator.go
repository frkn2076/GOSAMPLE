package nosql

import (
	"context"

	"app/GoSample/db"
	"app/GoSample/logger"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertLogRecord(record map[string]interface{}) {
	collection := db.MongoDB.Collection("RequestReponseLogs")
	_, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		logger.ErrorLog("An error occured while inserting log record to mongo db - Error:", err.Error(), "- LogRecord:", record)
	}
}

func GetLogRecord()  map[string]interface{} {
	var result map[string]interface{}
	collection := db.MongoDB.Collection("RequestReponseLogs")
	collection.FindOne(context.Background(), bson.D{}).Decode(&result)
	return result
}
