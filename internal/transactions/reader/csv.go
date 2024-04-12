package reader

import (
	"encoding/csv"
	"errors"
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"os"
)

var ErrFilePathRequired = errors.New("filepath is required")

type CSVReader struct{}

func NewCSVReader() CSVReader {
	return CSVReader{}
}

func (reader CSVReader) Read(options ...string) (domain.TransactionList, error) {
	if len(options) == 0 {
		return nil, ErrFilePathRequired
	}
	file, err := os.Open(options[0])

	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	transactions := domain.NewTransactionList()

	for index, rows := range records {
		if index == 0 {
			// skip titles
			continue
		}
		transaction, err := domain.NewTransaction(rows[0], rows[1], rows[2])
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}
