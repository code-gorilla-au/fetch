package fetch

import (
	"errors"
	"fmt"
)

var (
	ErrNoValidRetryStrategy = errors.New("no valid retry strategy")
)

type APIError struct {
	StatusCode int
	StatusText string
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: [%d]: %s", e.StatusText, e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return fmt.Errorf("%s: [%d]: %s", e.StatusText, e.StatusCode, e.Message)
}
