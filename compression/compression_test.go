package compression

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

func (suite *DriverSuite) TestDummy() {
	t := suite.T()
	assert.Equal(t, 1, 1, "Elementar equation is not equal")

}

func TestDriverSuite(t *testing.T) {
	suite.Run(t, new(DriverSuite))
}
