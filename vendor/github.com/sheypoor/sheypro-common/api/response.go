package api

import (
	"encoding/json"
	"net/http"
)

// Success writes a success response.
func Success(w http.ResponseWriter, result interface{}) {
	Response(w, http.StatusOK, result)
}

// Failure writes an error response.
func Failure(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *Error:
		reportError(w, err)
	case *ValidationError:
		reportValidationError(w, err)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Response writes a JSON response.
func Response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func reportError(w http.ResponseWriter, err error) {
	Response(w, err.(*Error).code, &Error{
		Message: err.(*Error).Message,
	})
}

func reportValidationError(w http.ResponseWriter, err error) {
	Response(w, http.StatusBadRequest, err)
}
