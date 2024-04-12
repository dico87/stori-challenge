package repository

import (
	"github.com/dico87/stori-challenge/internal/repository/entities"
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	db         *gorm.DB
	repository TransactionRepositoryImpl
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}

func initTestDB(showLog bool) (*gorm.DB, error) {
	if showLog {
		return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func (suite *TransactionRepositoryTestSuite) SetupTest() {
	db, err := initTestDB(true)
	suite.Nil(err)
	suite.db = db
	suite.repository = NewTransactionRepositoryImpl(suite.db)
	suite.MigrateTables()
}

func (suite *TransactionRepositoryTestSuite) MigrateTables() {
	suite.db.AutoMigrate(entities.Summary{})
	suite.db.AutoMigrate(entities.AverageTransaction{})
	suite.db.AutoMigrate(entities.GroupTransaction{})
	suite.db.AutoMigrate(entities.Transaction{})
}

func (suite *TransactionRepositoryTestSuite) Test_Create_Successfully() {
	// given
	transactions := domain.NewTransactionList()
	transaction, err := domain.NewTransaction("0", "7/15", "+60.5")
	suite.Nil(err)
	transactions = append(transactions, *transaction)
	transaction, err = domain.NewTransaction("1", "7/28", "-10.3")
	suite.Nil(err)
	transactions = append(transactions, *transaction)
	transaction, err = domain.NewTransaction("2", "8/2", "-20.46")
	suite.Nil(err)
	transactions = append(transactions, *transaction)
	transaction, err = domain.NewTransaction("3", "8/13", "+10")
	suite.Nil(err)
	transactions = append(transactions, *transaction)

	summary := domain.NewSummary(transactions.TotalBalance(), transactions.GroupByMonth(), transactions.Average())

	// when
	err = suite.repository.Create(summary, transactions)
	// then
	suite.Nil(err)

	var summaryEntity []entities.Summary
	suite.db.Find(&summaryEntity)
	suite.Len(summaryEntity, 1)

	var average []entities.AverageTransaction
	suite.db.Find(&average)
	suite.Len(average, 2)

	var group []entities.GroupTransaction
	suite.db.Find(&group)
	suite.Len(group, 2)

	var transactionEntity []entities.Transaction
	suite.db.Find(&transactionEntity)
	suite.Len(transactionEntity, 4)
}
