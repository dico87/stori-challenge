package main

import (
	"fmt"
	transactions2 "github.com/dico87/stori-challenge/internal/transactions"
	"github.com/dico87/stori-challenge/internal/transactions/reader"
	"github.com/dico87/stori-challenge/internal/transactions/sender"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("email_user")
	if user == "" {
		fmt.Printf("enviroment variable [%s] not found", "email_user")
		os.Exit(0)
	}
	password := os.Getenv("email_password")
	if password == "" {
		fmt.Printf("enviroment variable [%s] not found", "email_password")
		os.Exit(0)
	}
	smtpServer := os.Getenv("email_smtp_server")
	if smtpServer == "" {
		fmt.Printf("enviroment variable [%s] not found", "email_smtp_server")
		os.Exit(0)
	}
	to := os.Getenv("email_to")
	if to == "" {
		fmt.Printf("enviroment variable [%s] not found", "email_to")
		os.Exit(0)
	}
	subject := os.Getenv("email_subject")
	if subject == "" {
		fmt.Printf("enviroment variable [%s] not found", "email_subject")
		os.Exit(0)
	}

	toList := strings.Split(to, ",")

	csvReader := reader.NewCSVReader()
	emailSender := sender.NewEmailSender(user, password, smtpServer, user, toList, subject)

	transactions := transactions2.NewTransactions(emailSender, csvReader)
	err := transactions.Process()
	if err != nil {
		fmt.Printf("[ERROR] : %v", err)
		os.Exit(0)
	}

	fmt.Printf("!! Email send it !!")
}
