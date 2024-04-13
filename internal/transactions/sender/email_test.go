package sender

import (
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestEmailSenderTestSuite struct {
	suite.Suite
	emailSender Email
}

func TestEmailSenderSuite(t *testing.T) {
	suite.Run(t, new(TestEmailSenderTestSuite))
}

func (suite *TestEmailSenderTestSuite) SetupTest() {
	suite.emailSender = NewEmailSender("test@stori.com", "12345", "smtp.gmail.com", "test@story.com", []string{"test1@gmail.com"}, "Test Subject")
}

func (suite *TestEmailSenderTestSuite) Test_Send_Error_Send() {
	summary := domain.Summary{}
	err := suite.emailSender.Send(summary)
	suite.NotNil(err)
}
