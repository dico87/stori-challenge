package main

import (
	"fmt"
	"github.com/dico87/stori-challenge/internal/repository"
	"github.com/dico87/stori-challenge/internal/repository/entities"
	transactions2 "github.com/dico87/stori-challenge/internal/transactions"
	"github.com/dico87/stori-challenge/internal/transactions/reader"
	"github.com/dico87/stori-challenge/internal/transactions/sender"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
)

func main() {
	// Create email variables

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

	// Create Transactions Repository
	db, err := initDB()
	if err != nil {
		fmt.Printf("error init db [%v]", err)
		os.Exit(0)
	}

	repository := repository.NewTransactionRepositoryImpl(db)

	transactions := transactions2.NewTransactions(emailSender, csvReader, repository)
	err = transactions.Process()
	if err != nil {
		fmt.Printf("[ERROR] : %v", err)
		os.Exit(0)
	}

	fmt.Printf("!! Email send it !!")
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("transaction.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(entities.Summary{})
	db.AutoMigrate(entities.AverageTransaction{})
	db.AutoMigrate(entities.GroupTransaction{})
	db.AutoMigrate(entities.Transaction{})

	return db, nil
}
