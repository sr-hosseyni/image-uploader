package api

import (
	"regexp"
	"strings"
)

var (
	cellphoneRegexp  = regexp.MustCompile("^09\\d{9}")
	domainPartRegexp = regexp.MustCompile("^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])$")
)

// CheckCellphone checks validity of a cellphone number.
func CheckCellphone(cellphone string) bool {
	return cellphoneRegexp.MatchString(cellphone)
}

// NormalizeEmail computes a normalized representation of an email
// address. This method assumes the input as a valid email address.
func NormalizeEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		panic("invalid email address")
	}

	return parts[0] + "@" + strings.ToLower(parts[1])
}

// CheckEmail checks validity of an email address.
func CheckEmail(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	return CheckDomain(parts[1])
}

// CheckDomain checks validity of a domain address.
func CheckDomain(domain string) bool {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return false
	}

	for _, part := range parts {
		if !domainPartRegexp.MatchString(part) {
			return false
		}
	}

	return true
}
