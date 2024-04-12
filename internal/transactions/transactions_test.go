package transactions

import (
	"errors"
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"github.com/dico87/stori-challenge/internal/transactions/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionsTestSuite struct {
	suite.Suite
	sender       *mocks.SenderMock
	reader       *mocks.ReaderMock
	transactions Transactions
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsTestSuite))
}

func (suite *TransactionsTestSuite) SetupTest() {
	suite.sender = &mocks.SenderMock{}
	suite.reader = &mocks.ReaderMock{}
	suite.transactions = NewTransactions(suite.sender, suite.reader)
}

func (suite *TransactionsTestSuite) Test_Process_File_Error() {
	suite.reader.ExpectedRead = func(options ...string) (domain.TransactionList, error) {
		return nil, errors.New("error read file")
	}

	err := suite.transactions.Process()
	suite.Equal("error read file", err.Error())
}

func (suite *TransactionsTestSuite) Test_Process_Sender_Error() {
	suite.reader.ExpectedRead = func(options ...string) (domain.TransactionList, error) {
		return domain.NewTransactionList(), nil
	}

	suite.sender.ExpectedSend = func(summary domain.Summary) error {
		return errors.New("error send email")
	}

	err := suite.transactions.Process()
	suite.Equal("error send email", err.Error())
}

func (suite *TransactionsTestSuite) Test_Process_Successfully() {
	suite.reader.ExpectedRead = func(options ...string) (domain.TransactionList, error) {
		return domain.NewTransactionList(), nil
	}

	suite.sender.ExpectedSend = func(summary domain.Summary) error {
		return nil
	}

	err := suite.transactions.Process()
	suite.NoError(err)
}
