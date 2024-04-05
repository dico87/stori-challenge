package reader

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type CSVTestSuite struct {
	suite.Suite
	reader CSVReader
}

func TestCSVTestSuite(t *testing.T) {
	suite.Run(t, new(CSVTestSuite))
}

func (suite CSVTestSuite) SetupTest() {
	suite.reader = NewCSVReader()

}

func (suite CSVTestSuite) Test_FilePath_Empty_Options() {
	//given
	//when
	transactions, err := suite.reader.Read()
	//then
	suite.Nil(transactions)
	suite.Error(err)
	suite.Equal(ErrFilePathRequired, err)
}

func (suite CSVTestSuite) Test_FilePath_Wrong() {
	//given
	filePath := "$#@$%"
	//when
	transactions, err := suite.reader.Read(filePath)
	//then
	suite.Nil(transactions)
	suite.Error(err)
	suite.Equal(fmt.Sprintf("open %s: The system cannot find the file specified.", filePath), err.Error())
}

func (suite CSVTestSuite) Test_FilePath_Wrong_CSV_File() {
	//given
	current, err := os.Getwd()
	suite.NoError(err)
	filePath := fmt.Sprintf("%s/test_files/wrong.csv", current)
	//when
	transactions, err := suite.reader.Read(filePath)
	//then
	suite.Nil(transactions)
	suite.Error(err)
	suite.Equal("record on line 2: wrong number of fields", err.Error())
}

func (suite CSVTestSuite) Test_FilePath_Wrong_CSV_File_Data() {
	//given
	current, err := os.Getwd()
	suite.NoError(err)
	filePath := fmt.Sprintf("%s/test_files/wrong_data.csv", current)
	//when
	transactions, err := suite.reader.Read(filePath)
	//then
	suite.Nil(transactions)
	suite.Error(err)
	suite.Equal("not valid id", err.Error())
}

func (suite CSVTestSuite) Test_FilePath_Successfully() {
	//given
	current, err := os.Getwd()
	suite.NoError(err)
	filePath := fmt.Sprintf("%s/test_files/transactions.csv", current)
	//when
	transactions, err := suite.reader.Read(filePath)
	//then
	suite.NoError(err)
	suite.NotNil(transactions)
	suite.Equal(4, len(transactions))
}
