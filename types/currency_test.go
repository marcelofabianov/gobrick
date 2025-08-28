package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCurrency(t *testing.T) {
	testCases := []struct {
		name          string
		inputValue    string
		expected      Currency
		expectError   bool
		expectedError error
	}{
		{
			name:        "should create BRL currency successfully",
			inputValue:  "BRL",
			expected:    BRL,
			expectError: false,
		},
		{
			name:        "should create BRL currency successfully with lowercase input",
			inputValue:  "brl",
			expected:    BRL,
			expectError: false,
		},
		{
			name:        "should create USD currency successfully",
			inputValue:  "USD",
			expected:    USD,
			expectError: false,
		},
		{
			name:        "should create EUR currency successfully",
			inputValue:  "EUR",
			expected:    EUR,
			expectError: false,
		},
		{
			name:          "should return error for invalid currency",
			inputValue:    "GBP",
			expectError:   true,
			expectedError: ErrInvalidCurrency,
		},
		{
			name:          "should return error for empty string",
			inputValue:    "",
			expectError:   true,
			expectedError: ErrInvalidCurrency,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			currency, err := NewCurrency(tc.inputValue)

			if tc.expectError {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
				assert.Equal(t, Currency(""), currency)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, currency)
			}
		})
	}
}
