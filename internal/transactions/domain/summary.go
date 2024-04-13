package domain

import "fmt"

type Summary struct {
	Balance             float64
	GroupTransactions   map[string]int
	AverageTransactions map[string]float64
}

func NewSummary(balance float64, groupTransactions map[string]int, averageTransactions map[string]float64) Summary {
	return Summary{
		Balance:             balance,
		GroupTransactions:   groupTransactions,
		AverageTransactions: averageTransactions,
	}
}

func (s Summary) GroupTransactionsAsStringArray() []string {
	var transactions []string
	for date, countTransactions := range s.GroupTransactions {
		transactions = append(transactions, fmt.Sprintf("%s: %d\n", date, countTransactions))
	}

	return transactions
}

func (s Summary) AverageTransactionsAsStringArray() []string {
	var transactions []string
	for averageType, average := range s.AverageTransactions {
		transactions = append(transactions, fmt.Sprintf("%s: %.2f\n", averageType, average))
	}

	return transactions
}

func (s Summary) BalanceAsString() string {
	return fmt.Sprintf("Total balance is %.2f\n", s.Balance)
}

func (s Summary) AsString() string {
	stringTemplate := s.BalanceAsString()
	for _, transaction := range s.GroupTransactionsAsStringArray() {
		stringTemplate += transaction
	}

	for _, transaction := range s.AverageTransactionsAsStringArray() {
		stringTemplate += transaction
	}

	return stringTemplate
}
