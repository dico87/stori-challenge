package sender

import (
	"bytes"
	"github.com/dico87/stori-challenge/internal/transactions/domain"
	"html/template"
	"net/smtp"
)

const templatePath = "email-template.html"

type Email struct {
	user       string
	password   string
	smtpServer string
	from       string
	to         []string
	subject    string
}

func NewEmailSender(user string, password string, smtpServer string, from string, to []string, subject string) Email {
	return Email{
		user:       user,
		password:   password,
		smtpServer: smtpServer,
		from:       from,
		to:         to,
		subject:    subject,
	}
}

func (s Email) Send(summary domain.Summary) error {
	auth := smtp.PlainAuth("", s.user, s.password, s.smtpServer)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	emailSubject := "Subject: " + s.subject + "!\n"

	body, err := s.parseTemplate(templatePath, summary)
	if err != nil {
		return err
	}

	msg := []byte(emailSubject + mime + "\n" + *body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, s.from, s.to, msg); err != nil {
		return err
	}

	return nil
}

func (s Email) parseTemplate(templateFileName string, summary domain.Summary) (*string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, map[string]interface{}{
		"balance": summary.BalanceAsString(),
		"group":   summary.GroupTransactionsAsStringArray(),
		"average": summary.AverageTransactionsAsStringArray(),
	})

	if err != nil {
		return nil, err
	}

	template := buf.String()

	return &template, nil
}
