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

func (compressionDriver *CompressionDriver) AddString() error {

	entry := &bson.M{
		"_id": "compression",
		"counter": 3,
		"attributes": bson.M{
			"Romania": 0,
			"Bucharest": 1,
			"Azimut": 2,
		},
	}

	return compressionDriver.Connection.Insert(&entry)
}

func (compressionDriver *CompressionDriver) GetAttribute(findDict bson.M) ([]bson.M, error) {
	var results []bson.M
	err := compressionDriver.Connection.Find(findDict).All(&results)

	if err != nil {
		return nil, err
	}

	return results, nil
}
