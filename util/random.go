package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomly generate a random int64 number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min + 1) // 0 -> max-min
}

// randomly generated string of n characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i:=0; i<=n; i++{
		c := alphabet[(rand.Intn(k))]
		sb.WriteByte(c)
	}

	return sb.String()
}

// randomly generates a owner name
func RandomOwner() string {
	return RandomString(6)
}

// randomly generate money
func RandomMoney() int64{
	return RandomInt(0, 1000)
}

// randomly generats a currency code
func RandomCurrency() string {
	currencies := []string{EUR, USD, INR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// randomly generate an email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}