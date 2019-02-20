package util

import (
	"strings"
	"time"
)

const (
	hriDatePartFormat = "20060102"
	hriRandPartLength = 6
)

// GenerateHRI generates a unique human-readable ID.
func GenerateHRI(prefix string) string {
	if prefix == "" {
		prefix = "HRI"
	}

	datePart := time.Now().Format(hriDatePartFormat)
	randPart := RandomString(hriRandPartLength, CharsetUppercase)

	return strings.ToUpper(prefix) + "-" + datePart + "-" + randPart
}
