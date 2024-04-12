package mocks

import "github.com/dico87/stori-challenge/internal/transactions/domain"

type ReaderMock struct {
	ExpectedRead func(options ...string) (domain.TransactionList, error)
}

func (r *ReaderMock) Read(options ...string) (domain.TransactionList, error) {
	return r.ExpectedRead(options...)
}

type SenderMock struct {
	ExpectedSend func(summary domain.Summary) error
}

func (s *SenderMock) Send(summary domain.Summary) error {
	return s.ExpectedSend(summary)
}
