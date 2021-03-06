package compression

import (
	// "errors" errors.New("Mongopool streams config not found")
	"fmt"
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MONGO_HOST = "mongo"
const MONGO_PORT = 27017


type CompressionDriver struct {
	Connection *mgo.Collection
	KeyToValue map[string]map[string]int
	ValueToKey map[string]map[int]string
}

func NewDriver(database string, collection string) (*CompressionDriver, error) {
	host := fmt.Sprintf("%s:%d", MONGO_HOST, MONGO_PORT)
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	c := session.DB(database).C(collection)

	return &CompressionDriver{c, map[string]map[string]int{}, map[string]map[int]string{}}, nil
}

func (compressionDriver *CompressionDriver) AddString(entry *bson.M) error {
	return compressionDriver.Connection.Insert(&entry)
}

func (compressionDriver *CompressionDriver) DropCollection() {
	compressionDriver.Connection.DropCollection()
}

func (compressionDriver *CompressionDriver) UpdateKey(attribute string, key string) error {
	if _, ok := compressionDriver.KeyToValue[attribute][key]; ok {
		return nil
	}

	found := false
	counter := 0
	var entry bson.M
	findDict := bson.M{"_id": attribute}
	err := compressionDriver.Connection.Find(findDict).One(&entry)
	if err == nil {
		if _, ok := entry[key]; ok {
			found = true
			counter = entry[key].(int)
		} else {
			counter = entry["counter"].(int)
		}
	} else {
		counter = 0
		compressionDriver.KeyToValue[attribute] = map[string]int{}
		compressionDriver.ValueToKey[attribute] = map[int]string{}
	}
	if found == false {
		counter += 1
		updateDict := bson.M{"$set": bson.M{"counter": counter, key: counter}}
		compressionDriver.Connection.Update(findDict, updateDict)
	}

	// insert into local structure
	compressionDriver.KeyToValue[attribute][key] = counter
	compressionDriver.ValueToKey[attribute][counter] = key

	return nil
}

func (compressionDriver *CompressionDriver) UpdateValue(attribute string, value int) error {
	if _, ok := compressionDriver.ValueToKey[attribute][value]; ok == false {
		err := compressionDriver.LoadAttribute(attribute)
		if err != nil {
			return err
		}
	}
	if _, ok := compressionDriver.ValueToKey[attribute][value]; ok {
		return nil
	}
	return errors.New(fmt.Sprintf("Value %d does not exist for attribute %s", value, attribute))
}

func (compressionDriver *CompressionDriver) LoadAttribute(attribute string) error {
	var result bson.M
	findDict := bson.M{"_id": attribute}
	err := compressionDriver.Connection.Find(findDict).One(&result)

	if err != nil {
		return err
	}
	delete(result, "_id")
	delete(result, "counter")

	compressionDriver.KeyToValue[attribute] = convertKeyToValue(result)
	compressionDriver.ValueToKey[attribute] = reverseKeyToValue(result)

	return nil
}

func convertKeyToValue(object bson.M) map[string]int {
	result := map[string]int{}
	for key, value := range object {
		if valueInt, ok := value.(int); ok {
			result[key] = valueInt
		}
	}
	return result
}

func reverseKeyToValue(object bson.M) map[int]string {
	result := map[int]string{}
	for key, value := range object {
		if valueInt, ok := value.(int); ok {
			result[valueInt] = key
		}
	}
	return result
}