package accountModel

import (
	transactionModel "insan-service-account-backend/module/transaction/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	UserID uuid.UUID   `gorm:"unique;not null"` // Foreign key
	Transaction []transactionModel.Transaction `gorm:"foreignKey:AccountID"`
	AccountNumber string `json:"account_number"`
	Balance int `json:"balance"`
}
