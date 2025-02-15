package transactionModel

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Transaction struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	AccountID uuid.UUID `gorm:"not null" json:"account_id"` // ForeignKey
	TransactionType string `json:"transaction_type"`
	TransactionStatus string `json:"transaction_status"`
	Amount int `gorm:"not null" json:"amount"`
	TransactionDate time.Time `gorm:"not null" json:"transaction_date"`	
}
