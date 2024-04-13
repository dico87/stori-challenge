package repository

import (
	"github.com/dico87/stori-challenge/internal/repository/entities"
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepositoryImpl(db *gorm.DB) TransactionRepositoryImpl {
	return TransactionRepositoryImpl{
		db: db,
	}
}

func (t TransactionRepositoryImpl) Create(summary domain.Summary, transactions domain.TransactionList) error {
	db := t.db
	err := db.Transaction(func(tx *gorm.DB) error {
		// Create summary
		entitySummary := entities.Summary{}
		entitySummary.ToEntity(summary)

		tx.Create(&entitySummary)

		if tx.Error != nil {
			return tx.Error
		}

		// Create Group Transactions
		groupTransactions := entities.NewGroupTransactions(entitySummary.ID, summary)

		tx.CreateInBatches(&groupTransactions, 500)

		if tx.Error != nil {
			return tx.Error
		}

		// Create Average Transactions
		averageTransactions := entities.NewAverageTransactions(entitySummary.ID, summary)

		tx.CreateInBatches(&averageTransactions, 500)

		if tx.Error != nil {
			return tx.Error
		}

		// Create Transactions
		transactionsBatch := entities.NewTransactions(entitySummary.ID, transactions)

		tx.CreateInBatches(&transactionsBatch, 500)

		if tx.Error != nil {
			return tx.Error
		}

		return nil
	})

	return err
}
