package db

type IMongo interface {
	InsertRecord(collection string, record map[string]interface{})
}