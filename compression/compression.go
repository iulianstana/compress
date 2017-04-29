package compression

import (
	// "errors" errors.New("Mongopool streams config not found")
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MONGO_HOST = "mongo"
const MONGO_PORT = 27017


type CompressionDriver struct {
	Connection *mgo.Collection
	DATABASE   string
	COLLECTION string
}

func NewDriver(database string, collection string) (*CompressionDriver, error) {
	host := fmt.Sprintf("%s:%d", MONGO_HOST, MONGO_PORT)
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	c := session.DB(database).C(collection)

	return &CompressionDriver{c, database, collection}, nil
}

func (compressionDriver *CompressionDriver) AddString(entry *bson.M) error {
	return compressionDriver.Connection.Insert(&entry)
}

func (compressionDriver *CompressionDriver) GetAttribute(attribute string) ([]bson.M, error) {
	var results []bson.M
	findDict := bson.M{"_id": attribute}
	err := compressionDriver.Connection.Find(findDict).All(&results)

	if err != nil {
		return nil, err
	}

	return results, nil
}
