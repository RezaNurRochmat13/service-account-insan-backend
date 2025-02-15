package accountModel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	User User `json:"user"`
	AccountNumber string `json:"account_number"`
	Balance string `json:"balance"`
}
