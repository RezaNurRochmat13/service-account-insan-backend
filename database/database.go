package database

import (
	"fmt"
	"insan-service-account-backend/config"
	accountModel "insan-service-account-backend/module/account/model"
	transactionModel "insan-service-account-backend/module/transaction/model"
	userModel "insan-service-account-backend/module/user/model"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the database
func ConnectDatabase() {
	var err error

    p := config.Config("DB_PORT")
    port, err := strconv.ParseUint(p, 10, 32)

    if err != nil {
        log.Println("Could not parse port number")
    }
	
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the schema
	DB.AutoMigrate(&userModel.User{}, &accountModel.Account{}, &transactionModel.Transaction{})
}
