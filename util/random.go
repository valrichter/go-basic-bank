package util

import (
	"math"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

// Generates a random owner name
func RandomOwner() string {
	owner := gofakeit.FirstName() + " " + gofakeit.LastName()
	return owner
}

// Generates a random amount of money
func RandomMoney() float32 {
	rNum := gofakeit.Price(0, 1000)
	money := float32(math.Floor(float64(rNum*100)) / 100)
	return money
}

// Generates a random currency code
func RandomCurrency() string {
	currencies := []string{ARS, EUR, USD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// Generates a random email
func RandomEmail() string {
	email := gofakeit.Email()
	return email
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	n := min + rand.Int63n(max-min+1)
	return n
}

// RandomPassword generates a random password
func RandomPassword(length int) string {
	password := gofakeit.Password(false, false, false, false, false, length)
	return password
}

func RandomUsername() string {
	username := gofakeit.Username()
	return username
}
