package userModel

import (
	accountModel "insan-service-account-backend/module/account/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	Account accountModel.Account `gorm:"foreignKey:UserID"`
	Username string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
}