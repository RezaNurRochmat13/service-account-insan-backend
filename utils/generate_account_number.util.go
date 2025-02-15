package utils

import (
	"math/rand"
	"time"
)

func GenerateAccountNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	accountNumber := make([]byte, length)

	for i := range accountNumber {
		accountNumber[i] = digits[rand.Intn(len(digits))]
	}

	return "ACC-" + string(accountNumber)
}
