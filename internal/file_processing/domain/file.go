package domain

import (
	"errors"
	"github.com/dico87/stori-challenge/internal/platform/format"
	"strconv"
	"time"
)

var ErrIDNotValid = errors.New("id is not a number")

type Transaction struct {
	ID     uint
	Date   time.Time
	Amount string
}

func NewTransaction(id string, date string, amount string) (*Transaction, error) {
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrIDNotValid
	}

	dateAsTime, err := format.ParseDate(date)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		ID:     uint(idAsInt),
		Date:   *dateAsTime,
		Amount: amount,
	}, nil
}
