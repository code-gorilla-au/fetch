package fetch

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_Error(t *testing.T) {
	type fields struct {
		StatusCode int
		StatusText string
		Message    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return 4xx error",
			fields: fields{
				StatusCode: http.StatusBadRequest,
				StatusText: http.StatusText(http.StatusBadRequest),
				Message:    "there was an issue with the request",
			},
			want: "Bad Request [400]: there was an issue with the request",
		},
		{
			name: "should return 5xx error",
			fields: fields{
				StatusCode: http.StatusInternalServerError,
				StatusText: http.StatusText(http.StatusInternalServerError),
				Message:    "there was an issue with the request",
			},
			want: "Internal Server Error [500]: there was an issue with the request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := APIError{
				StatusCode: tt.fields.StatusCode,
				StatusText: tt.fields.StatusText,
				Message:    tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("APIError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIError_Unwrap_should_match(t *testing.T) {
	var err *APIError
	expectedErr := APIError{
		StatusCode: http.StatusBadRequest,
		StatusText: http.StatusText(http.StatusBadRequest),
		Message:    "some issue",
	}
	var testErr error = &expectedErr
	if errors.As(testErr, &err) {
		assert.Equal(t, err.Message, expectedErr.Message)
	} else {
		t.Error("non matching error")
	}
}

func TestAPIError_errors_is(t *testing.T) {
	fn := func() error {
		return &APIError{
			StatusCode: http.StatusOK,
		}
	}
	expected := APIError{
		StatusCode: http.StatusOK,
	}
	err := fn()
	assert.Equal(t, err.Error(), expected.Error())
}
