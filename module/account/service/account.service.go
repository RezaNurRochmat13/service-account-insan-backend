package accountService

import (
	"insan-service-account-backend/database"
	accountModel "insan-service-account-backend/module/account/model"
	userModel "insan-service-account-backend/module/user/model"
	"insan-service-account-backend/utils"

	"github.com/google/uuid"
)

func FindAllAccounts() ([]accountModel.Account, error) {
	db := database.DB
	var accounts []accountModel.Account

	if err := db.Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

func FindAccountByAccountNumber(accountNumber string) (*accountModel.Account, error) {
	db := database.DB
	var account accountModel.Account

	if err := db.Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func SaveAccounts(account *accountModel.Account) (*accountModel.Account, error) {
	db := database.DB

	// Check if the user exists
	var existingUser userModel.User

    if err := db.Where("id = ?", account.UserID).First(&existingUser).Error; err != nil {
        return nil, err
    }

	// Check if the account already exists
	var existingAccount accountModel.Account

	if err := db.Where("user_id = ?", account.UserID).First(&existingAccount).Error; err == nil {
		return nil, err
	}

    account.ID = uuid.New()
    account.AccountNumber = utils.GenerateAccountNumber(10)

	if err := db.Create(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}
