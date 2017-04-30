package compression

import (
	// "os"
	// "strings"
	"testing"

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

func (suite *DriverSuite) TearDownTest() {
	suite.driver.Connection.DropCollection()
}

func (suite *DriverSuite) TestInsertGet() {
	t := suite.T()

	attribute := "compression"
	entry := &bson.M{
		"_id": attribute,
		"counter": 3,
		"Romania": 0,
		"Bucharest": 1,
		"Azimut": 2,

	}

	// Insert into mongo
	err := suite.driver.AddString(entry)
	assert.Nil(t, err, "Entry was not Inserted")

	// Get from mongo
	err = suite.driver.LoadAttribute(attribute)
	assert.Nil(t, err, "Entry was not Inserted")
	assert.Equal(t, suite.driver.KeyToValue[attribute]["Romania"], 0, "Test attribute is returned")
}

func TestDriverSuite(t *testing.T) {
	suite.Run(t, new(DriverSuite))
}
