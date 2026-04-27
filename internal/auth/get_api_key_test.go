package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			expectedKey:   "valid_api_key",
			expectedError: nil,
		},
		{
			name:          "missing Authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InvalidHeader"},
			},
			expectedKey:   "",
			expectedError: assert.AnError, // Using assert.AnError to indicate any error is expected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			assert.Equal(t, tt.expectedKey, apiKey)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
