package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TransactionTestSuite struct {
	suite.Suite
	transactions TransactionList
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}

func (s *TransactionTestSuite) SetupTest() {
	s.transactions = NewTransactionList()
	transaction, err := NewTransaction("0", "7/15", "+60.5")
	s.Nil(err)
	s.transactions = append(s.transactions, *transaction)
	transaction, err = NewTransaction("1", "7/28", "-10.3")
	s.Nil(err)
	s.transactions = append(s.transactions, *transaction)
	transaction, err = NewTransaction("2", "8/2", "-20.46")
	s.Nil(err)
	s.transactions = append(s.transactions, *transaction)
	transaction, err = NewTransaction("3", "8/13", "+10")
	s.Nil(err)
	s.transactions = append(s.transactions, *transaction)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_ID() {
	// given
	id := "abc"
	date := "7/15"
	amount := "+60.5"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidID, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_Date_Less_Two_Fields() {
	// given
	id := "0"
	date := "2"
	amount := "+60.5"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidDate, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_Date_More_Three_Fields() {
	// given
	id := "0"
	date := "20/20/03/02"
	amount := "+60.5"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidDate, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_Date() {
	// given
	id := "0"
	date := "23/14"
	amount := "+60.5"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidDate, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Empty_Amount() {
	// given
	id := "0"
	date := "7/15"
	amount := "     "
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrEmptyAmount, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_Amount_Less_Two_Characters() {
	// given
	id := "0"
	date := "7/15"
	amount := "*"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidAmount, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_Amount_With_Invalid_Sign() {
	// given
	id := "0"
	date := "7/15"
	amount := "*30.4"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidAmount, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_With_Not_Valid_Amount_Not_Number() {
	// given
	id := "0"
	date := "7/15"
	amount := "+abc"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(transaction)
	s.Equal(ErrNotValidAmount, err)
}

func (s *TransactionTestSuite) Test_NewTransaction_Successfully() {
	// given
	id := "0"
	date := "7/15/2024"
	amount := "+60.15"
	// when
	transaction, err := NewTransaction(id, date, amount)
	// then
	s.Nil(err)
	s.Equal(uint(0), transaction.ID)
	s.Equal("2024-07-15", transaction.Date.Format(time.DateOnly))
	s.Equal("+60.15", transaction.Amount)
}

func (s *TransactionTestSuite) Test_TotalBalance() {
	// given
	// when
	balance := s.transactions.TotalBalance()
	// then
	s.Equal(39.74, balance)
}

func (s *TransactionTestSuite) Test_GroupByMonth() {
	// given
	// when
	group := s.transactions.GroupByMonth()
	// then
	s.Equal(2, group["July of 2024"])
	s.Equal(2, group["August of 2024"])
}

func (s *TransactionTestSuite) Test_Average() {
	// given
	// when
	group := s.transactions.Average()
	// then
	s.Equal(-15.38, group["debit"])
	s.Equal(35.25, group["credit"])
}
