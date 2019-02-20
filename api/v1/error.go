package v1

import (
	"net/http"

	"github.com/sheypoor/sheypro-common/api"
)

var (
	ErrInvalidJSON   = api.NewError("error.invalid-json", http.StatusBadRequest)
	ErrCreateAdmin   = api.NewError("error.create-admin", http.StatusForbidden)
	ErrUserExists    = api.NewError("error.user-exists", http.StatusConflict)
	ErrInvalidClient = api.NewError("error.invalid-client", http.StatusUnauthorized)
	ErrUnauthorized  = api.NewError("error.unauthorized", http.StatusUnauthorized)
	ErrUserNotFound  = api.NewError("error.user-not-found", http.StatusNotFound)
	ErrUserDeleted   = api.NewError("error.user-deleted", http.StatusGone)

	ErrUnsupportedGrant    = api.NewError("error.unsupported-grant-type", http.StatusBadRequest)
	ErrInvalidCredentials  = api.NewError("error.invalid-credentials", http.StatusUnauthorized)
	ErrInvalidAccessToken  = api.NewError("error.invalid-access-token", http.StatusUnauthorized)
	ErrExpiredAccessToken  = api.NewError("error.expired-access-token", http.StatusUnauthorized)
	ErrInvalidRefreshToken = api.NewError("error.invalid-refresh-token", http.StatusUnauthorized)
	ErrExpiredRefreshToken = api.NewError("error.expired-refresh-token", http.StatusUnauthorized)
)
