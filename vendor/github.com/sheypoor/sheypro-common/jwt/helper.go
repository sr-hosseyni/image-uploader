package jwt

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AccessTokenLifeTime         = 3600
	AccessTokenLifeTimeDuration = AccessTokenLifeTime * time.Second

	AdminUserType      = "admin"
	RealEstateUserType = "real-estate"
	CarSaleUserType    = "car-sale"

	BasicScope = "basic"
	FullScope  = "full"
)

var (
	// ErrInvalidAccessToken reports an invalid access token.
	ErrInvalidAccessToken = errors.New("access token is invalid")

	// ErrExpiredAccessToken reports an expired access token.
	ErrExpiredAccessToken = errors.New("access token is expired")

	// ErrInvalidRefreshToken reports an invalid refresh token.
	ErrInvalidRefreshToken = errors.New("refresh token is invalid")
)

// AccessTokenSubject is the payload which is injected into access
// tokens as their subject.
type AccessTokenSubject struct {
	User     string `json:"user"`
	UserType string `json:"user_type"`
	Client   string `json:"client"`
	Scope    string `json:"scope"`
}

// RefreshTokenSubject is the payload which is injected into refresh
// tokens as their subject.
type RefreshTokenSubject struct {
	User   string `json:"user"`
	Client string `json:"client"`
}

// IssueAccessToken creates a new access token.
func IssueAccessToken(user, userType, client, scope, key string) (string, error) {
	bytes, err := json.Marshal(&AccessTokenSubject{user, userType, client, scope})
	if err != nil {
		return "", err
	}

	now := time.Now()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   string(bytes),
		ExpiresAt: now.Add(AccessTokenLifeTimeDuration).Unix(),
		IssuedAt:  now.Unix(),
	})

	return accessToken.SignedString([]byte(key))
}

// IssueRefreshToken creates a new refresh token.
func IssueRefreshToken(user, client, key string) (string, error) {
	bytes, err := json.Marshal(&RefreshTokenSubject{user, client})
	if err != nil {
		return "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:  string(bytes),
		IssuedAt: time.Now().Unix(),
	})

	return refreshToken.SignedString([]byte(key))
}

// ParseAccessToken decodes an access token and extracts its subject.
func ParseAccessToken(accessToken, key string) (*AccessTokenSubject, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidAccessToken
		}
		return []byte(key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, ErrExpiredAccessToken
		}
		return nil, ErrInvalidAccessToken
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, ErrInvalidAccessToken
	}

	subject := &AccessTokenSubject{}
	if err := json.Unmarshal([]byte(claims.Subject), subject); err != nil {
		return nil, ErrInvalidAccessToken
	}

	return subject, nil
}

// ParseRefreshToken decodes a refresh token and extracts its subject.
func ParseRefreshToken(refreshToken, key string) (*RefreshTokenSubject, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidRefreshToken
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, ErrInvalidRefreshToken
	}

	subject := &RefreshTokenSubject{}
	if err := json.Unmarshal([]byte(claims.Subject), subject); err != nil {
		return nil, ErrInvalidRefreshToken
	}

	return subject, nil
}
