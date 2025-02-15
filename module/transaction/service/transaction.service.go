package transactionService

import (
	"errors"
	"insan-service-account-backend/database"
	accountModel "insan-service-account-backend/module/account/model"
	transactionModel "insan-service-account-backend/module/transaction/model"
	"time"

	"github.com/google/uuid"
)

func FindAllTransactions() ([]transactionModel.Transaction, error) {
	db := database.DB
	var transactions []transactionModel.Transaction

	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func FindTransactionByAccountNumber(accountID string) ([]transactionModel.Transaction, error) {
	db := database.DB
	var transactions []transactionModel.Transaction

	if err := db.Where("account_id = ?", accountID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func CreateNewTransaction(transaction *transactionModel.Transaction) (*transactionModel.Transaction, error) {
	db := database.DB

	// Check account exists
	var checkAccount accountModel.Account

	if err := db.Where("id = ?", transaction.AccountID).First(&checkAccount).Error; err != nil {
		return nil, err
	}

	// Check transaction type
	if transaction.TransactionType == "DEPOSIT" {
		transaction.TransactionStatus = "SUCESSS"
		checkAccount.Balance += transaction.Amount
	} 
	
	if transaction.TransactionType == "WITHDRAW" {
		if transaction.Amount > checkAccount.Balance {
			return nil, errors.New("insufficient balance")
		}
		transaction.TransactionStatus = "SUCESSS"
		checkAccount.Balance -= transaction.Amount
	}

	// Generate UUID and set timestamp
	transaction.ID = uuid.New()
	transaction.TransactionDate = time.Now()

	// Save the transaction
	if err := db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	// Update account balance in DB
	if err := db.Save(&checkAccount).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

