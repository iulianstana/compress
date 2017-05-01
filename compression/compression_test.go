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

func (suite *DriverSuite) TestUpdateValue() {
	t := suite.T()
	attribute := "compression"
	entry := bson.M{
		"_id": attribute,
		"counter": 3,
		"Romania": 0,
		"Bucharest": 1,
		"Azimut": 2,

	}
	err := suite.driver.Connection.Insert(&entry)
	assert.Nil(t, err, "Mongo could not Insert entry")

	// Value is already in ValueToKey structure
	err = suite.driver.UpdateValue(attribute, 0)
	assert.Nil(t, err, "Value was not fount in ValueToKey Structure")

	// Reload structure and test if my value appeared
	delete(suite.driver.ValueToKey[attribute], 0)
	err = suite.driver.UpdateValue(attribute, 0)
	assert.Nil(t, err, "Value was not fount in ValueToKey Structure")

	// Value is not found in ValueToKey structure return err
	expectReturnError := errors.New(fmt.Sprintf("Value 10 does not exist for attribute %s", attribute))
	err = suite.driver.UpdateValue(attribute, 10)
	assert.Equal(t, err, expectReturnError, "Error message was not returned")
}


func (suite *DriverSuite) TestLoadAttribute() {
	t := suite.T()

	attribute := "compression"
	entry := bson.M{
		"_id": attribute,
		"counter": 3,
		"Romania": 0,
		"Bucharest": 1,
		"Azimut": 2,

	}
	err := suite.driver.Connection.Insert(&entry)
	assert.Nil(t, err, "Mongo could not Insert entry")


	// Load attribute in our driver
	err = suite.driver.LoadAttribute(attribute)
	err = suite.driver.LoadAttribute(attribute)

	assert.Nil(t, err, "Mongo could not find any result")
	assert.Equal(t, len(suite.driver.KeyToValue[attribute]), len(entry) - 2, "There are different number of elements")
	assert.Equal(t, len(suite.driver.ValueToKey[attribute]), len(entry) - 2, "There are different number of elements")
	assert.Equal(t,suite.driver.ValueToKey[attribute][0], "Romania", "Value number has other key")
	assert.Equal(t,suite.driver.KeyToValue[attribute]["Romania"], 0, "Key element has other value")
}

func (suite *DriverSuite) TearDownTest() {
	suite.driver.Connection.DropCollection()
}

func TestDriverSuite(t *testing.T) {
	suite.Run(t, new(DriverSuite))
}
