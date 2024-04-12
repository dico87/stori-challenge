package transactions

import (
	"github.com/dico87/stori-challenge/internal/transactions/domain"
)

type Transactions struct {
	sender Sender
	reader Reader
}

type Sender interface {
	Send(summary domain.Summary) error
}

type Reader interface {
	Read(options ...string) (domain.TransactionList, error)
}

func NewTransactions(sender Sender, reader Reader) Transactions {
	return Transactions{
		sender: sender,
		reader: reader,
	}
}

func (t Transactions) Process() error {
	transactions, err := t.reader.Read("sample.csv")
	if err != nil {
		return err
	}

	summary := domain.NewSummary(transactions.TotalBalance(), transactions.GroupByMonth(), transactions.Average())

	err = t.sender.Send(summary)
	if err != nil {
		return err
	}

	return nil
}
