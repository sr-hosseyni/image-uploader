package api

// API Error ---------------------------------------------------------

// NewError instantiates a new Error.
func NewError(message string, code int) error {
	return &Error{message, code}
}

// Error represents an API error.
type Error struct {
	Message string `json:"error"`
	code    int
}

// Code retrieves error's code.
func (e *Error) Code() int {
	return e.code
}

// Error implements error.Error.
func (e *Error) Error() string {
	return e.Message
}

// Validation Error --------------------------------------------------

// NewValidationError instantiates a new ValidationError.
func NewValidationError(items ...SingleValidationError) error {
	return &ValidationError{items}
}

// SingleValidationError represents a validation error for a single
// input field.
type SingleValidationError struct {
	Source string `json:"source"`
	Reason string `json:"reason"`
}

// Error implements error.Error.
func (e *SingleValidationError) Error() string {
	return e.Reason
}

// ValidationError represents a validation error for an input form.
type ValidationError struct {
	Errors []SingleValidationError `json:"errors"`
}

// Error implements error.Error.
func (e *ValidationError) Error() string {
	return "validation error"
}
