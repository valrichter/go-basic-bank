package util

import (
	"math"
	"math/rand"
	"strings"

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

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
