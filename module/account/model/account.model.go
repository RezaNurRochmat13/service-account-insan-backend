package accountModel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	UserID uuid.UUID   `gorm:"unique;not null"` // Foreign key
	AccountNumber string `json:"account_number"`
	Balance int `json:"balance"`
}
