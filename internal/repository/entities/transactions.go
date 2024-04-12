package entities

import (
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"time"
)

type Summary struct {
	ID      uint    `gorm:"primaryKey;column:id"`
	Balance float64 `gorm:"column:balance"`
}

func (*Summary) TableName() string {
	return "summary"
}

func (s *Summary) ToEntity(summary domain.Summary) {
	s.Balance = summary.Balance
}

type Transaction struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	SummaryID uint      `gorm:"column:summary_id"`
	Date      time.Time `gorm:"column:date"`
	Amount    string    `gorm:"column:amount"`
}

func (Transaction) TableName() string {
	return "transaction"
}

func newTransaction(summaryID uint, transaction domain.Transaction) Transaction {
	return Transaction{
		SummaryID: summaryID,
		Date:      transaction.Date,
		Amount:    transaction.Amount,
	}
}

type Transactions []Transaction

func NewTransactions(summaryID uint, list domain.TransactionList) Transactions {
	var transactions []Transaction
	for _, transaction := range list {
		transactions = append(transactions, newTransaction(summaryID, transaction))
	}

	return transactions
}

type GroupTransaction struct {
	ID        uint   `gorm:"primaryKey;column:id"`
	SummaryID uint   `gorm:"column:summary_id"`
	Date      string `gorm:"column:date"`
	Count     int    `gorm:"column:count"`
}

func (*GroupTransaction) TableName() string {
	return "group_transaction"
}

type GroupTransactions []GroupTransaction

func NewGroupTransactions(summaryID uint, summary domain.Summary) GroupTransactions {
	var groupTransactions []GroupTransaction
	for group, count := range summary.GroupTransactions {
		groupTransactions = append(groupTransactions, GroupTransaction{SummaryID: summaryID, Date: group, Count: count})
	}

	return groupTransactions
}

type AverageTransaction struct {
	ID        uint    `gorm:"primaryKey;column:id"`
	SummaryID uint    `gorm:"column:summary_id"`
	Type      string  `gorm:"column:type"`
	Amount    float64 `gorm:"column:amount"`
}

func (AverageTransaction) TableName() string {
	return "average_transaction"
}

type AverageTransactions []AverageTransaction

func NewAverageTransactions(summaryID uint, summary domain.Summary) AverageTransactions {
	var averageTransactions []AverageTransaction
	for transactionType, average := range summary.AverageTransactions {
		averageTransactions = append(averageTransactions, AverageTransaction{SummaryID: summaryID, Type: transactionType, Amount: average})
	}

	return averageTransactions
}
