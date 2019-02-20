package util

import (
	"math/rand"
	"time"
)

const (
	CharsetAll       = CharsetUppercase | CharsetLowercase | CharsetDigit
	CharsetUppercase = 1
	CharsetLowercase = 2
	CharsetDigit     = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string.
func RandomString(n, charset int) string {
	if n <= 0 {
		panic("invalid random string length")
	}

	chars := []byte(buildCharset(charset))
	charsLength := len(chars)
	if charsLength == 0 {
		panic("charset is empty")
	}

	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, chars[rand.Intn(charsLength)])
	}

	return string(result)
}

func buildCharset(charset int) string {
	chars := ""
	if charset&CharsetUppercase != 0 {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if charset&CharsetLowercase != 0 {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if charset&CharsetDigit != 0 {
		chars += "0123456789"
	}

	return chars
}
