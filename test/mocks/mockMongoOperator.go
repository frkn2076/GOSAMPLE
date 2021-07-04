package mocks

type MockMongoOperator struct{}

func (u MockMongoOperator) InsertRecord(collectionName string, record map[string]interface{}) {}