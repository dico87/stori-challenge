package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SummaryTestSuite struct {
	suite.Suite
	summary Summary
}

func TestSummarySuite(t *testing.T) {
	suite.Run(t, new(SummaryTestSuite))
}

func (suite *SummaryTestSuite) SetupTest() {
	balance := 39.74
	groupTransactions := map[string]int{
		"Number of transactions in July":    2,
		"Number of transactions in August:": 2,
	}

	averageTransactions := map[string]float64{
		"Average debit amount":  -15.38,
		"Average credit amount": 35.25,
	}
	suite.summary = NewSummary(balance, groupTransactions, averageTransactions)
}

func (suite *SummaryTestSuite) Test_AsString() {
	//given
	//when
	asString := suite.summary.AsString()
	//then

	suite.Equal("Total balance is 39.74\nNumber of transactions in July: 2\nNumber of transactions in August:: 2\nAverage debit amount: -15.38\nAverage credit amount: 35.25\n", asString)
}
