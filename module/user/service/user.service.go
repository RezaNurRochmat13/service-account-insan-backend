package userService

import (
	"errors"
	"insan-service-account-backend/database"
	userModel "insan-service-account-backend/module/user/model"

	"github.com/google/uuid"
)


func FindAllUsers() ([]userModel.User, error) {
	db := database.DB
	var users []userModel.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUser(user *userModel.User) (*userModel.User, error) {
	db := database.DB

	var existingUser userModel.User

	if err := db.Where("phone_number = ?", user.PhoneNumber).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this phone number already exists")
	}

	// Generate a new UUID for the user
	user.ID = uuid.New()

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
