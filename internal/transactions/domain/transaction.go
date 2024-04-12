package domain

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	debit         = "debit"
	credit        = "credit"
	decimals      = 2
	dateSeparator = "/"
)

var (
	ErrNotValidDate   = errors.New("not valid date %s")
	ErrNotValidID     = errors.New("not valid id")
	ErrNotValidAmount = errors.New("not valid amount")
	ErrEmptyAmount    = errors.New("amount is empty")
)

type Transaction struct {
	ID     uint
	Date   time.Time
	Amount string
}

type TransactionList []Transaction

func NewTransaction(id string, date string, amount string) (*Transaction, error) {
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrNotValidID
	}

	dateAsTime, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(amount, " ")) == 0 {
		return nil, ErrEmptyAmount
	}

	if len(amount) < 2 {
		return nil, ErrNotValidAmount
	}

	if amount[:1] != "-" && amount[:1] != "+" {
		return nil, ErrNotValidAmount
	}

	_, err = strconv.ParseFloat(amount[1:], 64)
	if err != nil {
		return nil, ErrNotValidAmount
	}

	return &Transaction{
		ID:     uint(idAsInt),
		Date:   *dateAsTime,
		Amount: amount,
	}, nil
}

func NewTransactionList() TransactionList {
	return make(TransactionList, 0)
}

func (transaction Transaction) IsDebit() bool {
	return transaction.Amount[:1] == "-"
}

func (transaction Transaction) AmountValue() float64 {
	number, _ := strconv.ParseFloat(transaction.Amount[1:], 64)
	return number
}

func (transactions TransactionList) TotalBalance() float64 {
	balance := 0.0

	for _, transaction := range transactions {
		if transaction.IsDebit() {
			balance -= transaction.AmountValue()
			continue
		}
		balance += transaction.AmountValue()
	}

	return balance
}

func (transactions TransactionList) GroupByMonth() map[string]int {
	group := make(map[string]int)

	for _, transaction := range transactions {
		key := fmt.Sprintf("%s of %d", transaction.Date.Month().String(), transaction.Date.Year())
		if _, ok := group[key]; !ok {
			group[key] = 0
		}
		group[key] = group[key] + 1
	}

	return group
}

func (transactions TransactionList) Average() map[string]float64 {
	sum := map[string]float64{debit: 0.0, credit: 0.0}
	count := map[string]int{debit: 0, credit: 0}

	for _, transaction := range transactions {
		if transaction.IsDebit() {
			sum[debit] = sum[debit] - transaction.AmountValue()
			count[debit] = count[debit] + 1
			continue
		}
		sum[credit] = sum[credit] + transaction.AmountValue()
		count[credit] = count[credit] + 1
	}

	return map[string]float64{
		debit:  roundFloat(sum[debit]/float64(count[debit]), decimals),
		credit: roundFloat(sum[credit]/float64(count[credit]), decimals),
	}
}

func parseDate(date string) (*time.Time, error) {
	fields := strings.Split(date, dateSeparator)
	if len(fields) < 2 {
		fmt.Println(date)
		return nil, ErrNotValidDate
	}

	if len(fields) > 3 {
		fmt.Println(date)
		return nil, ErrNotValidDate
	}

	dateAsString := fmt.Sprintf("%d-%s-%s", time.Now().Year(), fmt.Sprintf("%02s", fields[0]), fmt.Sprintf("%02s", fields[1]))
	if len(fields) == 3 {
		dateAsString = fmt.Sprintf("%s-%s-%s", fields[2], fmt.Sprintf("%02s", fields[0]), fmt.Sprintf("%02s", fields[1]))
	}

	dateAsTime, err := time.Parse(time.DateOnly, dateAsString)
	if err != nil {
		fmt.Println(dateAsString)
		return nil, ErrNotValidDate
	}

	return &dateAsTime, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
