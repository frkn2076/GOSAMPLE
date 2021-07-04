package db

import(
	"context"

	"app/GoSample/logger"
)

var Mongo *MongoOperator

func init() {
	Mongo = new(MongoOperator)
}

type MongoOperator struct{}

func (u *MongoOperator) InsertRecord(collectionName string, record map[string]interface{}) {
	collection := MongoDB.Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), record)
	if err != nil {
		logger.ErrorLog("An error occured while inserting log record to mongo db - Error:", err.Error(), "- LogRecord:", record)
	}
}