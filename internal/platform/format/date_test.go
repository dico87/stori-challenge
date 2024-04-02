package format

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type DateTestSuite struct {
	suite.Suite
}

func TestDateSuite(t *testing.T) {
	suite.Run(t, new(DateTestSuite))
}

func (suite DateTestSuite) Test_ParseFormat_Not_Valid_Date_Less_Two_Fields() {
	// given
	dateAsString := "05"
	// when
	date, err := ParseDate(dateAsString)
	// then
	suite.Nil(date)
	suite.Error(err)
	suite.Equal(ErrNotValidFormat, err)
}

func (suite DateTestSuite) Test_ParseFormat_Not_Valid_Date_More_Three_Fields() {
	// given
	dateAsString := "05/05/2003/12"
	// when
	date, err := ParseDate(dateAsString)
	// then
	suite.Nil(date)
	suite.Error(err)
	suite.Equal(ErrNotValidFormat, err)
}

func (suite DateTestSuite) Test_ParseFormat_Not_Valid_Date() {
	// given
	dateAsString := "15/05/2003"
	// when
	date, err := ParseDate(dateAsString)
	// then
	suite.Nil(date)
	suite.Error(err)
	suite.Equal(ErrNotValidFormat, err)
}

func (suite DateTestSuite) Test_ParseFormat_Successfully_With_Two_Fields() {
	// given
	dateAsString := "7/28"
	// when
	date, err := ParseDate(dateAsString)
	// then
	suite.NoError(err)
	suite.Equal("July", date.Month().String())
	suite.Equal(28, date.Day())
	suite.Equal(time.Now().Year(), date.Year())
}

func (suite DateTestSuite) Test_ParseFormat_Successfully_With_Three_Fields() {
	// given
	dateAsString := "7/28/2024"
	// when
	date, err := ParseDate(dateAsString)
	// then
	suite.NoError(err)
	suite.Equal("July", date.Month().String())
	suite.Equal(28, date.Day())
	suite.Equal(2024, date.Year())
}
