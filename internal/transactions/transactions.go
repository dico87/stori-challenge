package transactions

import (
	"github.com/dico87/stori-challenge/internal/transactions/domain"
)

type Sender interface {
	Send(summary domain.Summary) error
}

type Reader interface {
	Read(options ...string) (domain.TransactionList, error)
}

type TransactionRepository interface {
	Create(summary domain.Summary, transactions domain.TransactionList) error
}

type Transactions struct {
	sender     Sender
	reader     Reader
	repository TransactionRepository
}

func NewTransactions(sender Sender, reader Reader, repository TransactionRepository) Transactions {
	return Transactions{
		sender:     sender,
		reader:     reader,
		repository: repository,
	}
}

func (t Transactions) Process() error {
	transactions, err := t.reader.Read("large_sample.csv")

	if err != nil {
		return err
	}

	summary := domain.NewSummary(transactions.TotalBalance(), transactions.GroupByMonth(), transactions.Average())

	err = t.repository.Create(summary, transactions)

	if err != nil {
		return err
	}

	err = t.sender.Send(summary)
	if err != nil {
		return err
	}

	return nil
}
