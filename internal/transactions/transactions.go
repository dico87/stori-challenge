package transactions

import (
	"fmt"
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"time"
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
	startTime := time.Now()
	transactions, err := t.reader.Read("large_sample.csv")
	endTime := time.Now()
	fmt.Printf("Time read : [%d]", endTime.Sub(startTime).Milliseconds())
	if err != nil {
		return err
	}

	startTime = time.Now()
	summary := domain.NewSummary(transactions.TotalBalance(), transactions.GroupByMonth(), transactions.Average())
	endTime = time.Now()
	fmt.Printf("Time build summary : [%d]", endTime.Sub(startTime).Milliseconds())

	startTime = time.Now()
	err = t.repository.Create(summary, transactions)
	endTime = time.Now()
	fmt.Printf("Time save in database : [%d]", endTime.Sub(startTime).Milliseconds())

	if err != nil {
		return err
	}

	err = t.sender.Send(summary)
	if err != nil {
		return err
	}

	return nil
}
