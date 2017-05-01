package compression

import (
	// "os"
	// "strings"
	"testing"
	"fmt"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"gopkg.in/mgo.v2/bson"
)

const DATABASE = "test"
const COLLECTION = "compression"

type DriverSuite struct {
	driver *CompressionDriver
	suite.Suite
}

func (suite *DriverSuite) SetupSuite() {
	suite.driver, _ = NewDriver(DATABASE, COLLECTION)
}

func (suite *DriverSuite) TestUpdateKey() {
	t := suite.T()

	attribute, entry := generateEntry()
	err := suite.driver.Connection.Insert(&entry)
	assert.Nil(t, err, "Mongo could not Insert entry")

	entryKey := "Romania"
	// Key is already in KeyToValue structure
	mapSize := len(suite.driver.KeyToValue[attribute])
	err = suite.driver.UpdateKey(attribute, entryKey)
	assert.Nil(t, err, "Key was not found in KeyToValue Structure")
	assert.Equal(t, mapSize, len(suite.driver.KeyToValue[attribute]), "Structure was Updated")

	// Key appeared in the mean time
	delete(suite.driver.KeyToValue[attribute], entryKey)

	mapSize = len(suite.driver.KeyToValue[attribute])
	err = suite.driver.UpdateKey(attribute, entryKey)
	assert.Nil(t, err, "Key was not fount in KeyToValue Structure")
	assert.Contains(t, suite.driver.KeyToValue[attribute], entryKey, "Key was not inserted")
	assert.Equal(t, mapSize + 1, len(suite.driver.KeyToValue[attribute]), "Key was not updated")

	// Key not present, insert it
	mapSize = len(suite.driver.KeyToValue[attribute])
	newKey := "GoLang"
	err = suite.driver.UpdateKey(attribute, newKey)
	assert.Nil(t, err, "Value was not fount in ValueToKey Structure")
	assert.Contains(t, suite.driver.KeyToValue[attribute], newKey, "Key was not inserted")
	assert.Equal(t, mapSize + 1, len(suite.driver.KeyToValue[attribute]), "New Key was not added")

	// Insert Kew in new attribute
	newAttribute := "NewAttribute"
	err = suite.driver.UpdateKey(newAttribute, newKey)
	assert.Nil(t, err, "Value was not fount in ValueToKey Structure")
	assert.Contains(t, suite.driver.KeyToValue, newAttribute, "NewAttribute was not inserted")
	assert.Equal(t, 1, len(suite.driver.KeyToValue[newAttribute]), "NewAttribute was not created")
}

func (suite *DriverSuite) TestUpdateValue() {
	t := suite.T()

	attribute, entry := generateEntry()
	err := suite.driver.Connection.Insert(&entry)
	assert.Nil(t, err, "Mongo could not Insert entry")

	entryValue := 1
	// Value is already in ValueToKey structure
	err = suite.driver.UpdateValue(attribute, entryValue)
	assert.Nil(t, err, "Value was not fount in ValueToKey Structure")

	// Reload structure and test if my value appeared
	delete(suite.driver.ValueToKey[attribute], entryValue)
	err = suite.driver.UpdateValue(attribute, entryValue)
	assert.Nil(t, err, "Value was not fount in ValueToKey Structure")

	// Value is not found in ValueToKey structure return err
	unknownValue := entry["counter"].(int) + 10
	expectReturnError := errors.New(fmt.Sprintf("Value %d does not exist for attribute %s", unknownValue, attribute))
	err = suite.driver.UpdateValue(attribute, unknownValue)
	assert.Equal(t, err, expectReturnError, "Error message was not returned")
}


func (suite *DriverSuite) TestLoadAttribute() {
	t := suite.T()

	attribute, entry := generateEntry()
	err := suite.driver.Connection.Insert(&entry)
	assert.Nil(t, err, "Mongo could not Insert entry")

	key, value := "Romania", 1 // use random element from entry
	// Load attribute in our driver
	err = suite.driver.LoadAttribute(attribute)
	err = suite.driver.LoadAttribute(attribute)

	assert.Nil(t, err, "Mongo could not find any result")
	assert.Equal(t, len(suite.driver.KeyToValue[attribute]), len(entry) - 2, "There are different number of elements")
	assert.Equal(t, len(suite.driver.ValueToKey[attribute]), len(entry) - 2, "There are different number of elements")
	assert.Equal(t,suite.driver.ValueToKey[attribute][value], key, "Value number has other key")
	assert.Equal(t,suite.driver.KeyToValue[attribute][key], value, "Key element has other value")
}

func generateEntry() (string, bson.M) {
	attribute := "compression"
	entry := bson.M{
		"_id": attribute,
		"counter": 3,
		"Romania": 1,
		"Bucharest": 2,
		"Azimut": 3,

	}
	return attribute, entry
}
func (suite *DriverSuite) TearDownTest() {
	suite.driver.Connection.DropCollection()
}

func TestDriverSuite(t *testing.T) {
	suite.Run(t, new(DriverSuite))
}
