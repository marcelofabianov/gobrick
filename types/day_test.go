package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/marcelofabianov/gobrick/types"
)

func TestNewDay(t *testing.T) {
	testCases := []struct {
		name          string
		inputValue    int
		expectedDay   types.Day
		expectError   bool
		expectedError error
	}{
		{
			name:        "should create day successfully with lower bound",
			inputValue:  1,
			expectedDay: types.Day(1),
			expectError: false,
		},
		{
			name:        "should create day successfully with upper bound",
			inputValue:  31,
			expectedDay: types.Day(31),
			expectError: false,
		},
		{
			name:        "should create day successfully with middle value",
			inputValue:  15,
			expectedDay: types.Day(15),
			expectError: false,
		},
		{
			name:          "should return error for value below lower bound",
			inputValue:    0,
			expectError:   true,
			expectedError: types.ErrInvalidDay,
		},
		{
			name:          "should return error for value above upper bound",
			inputValue:    32,
			expectError:   true,
			expectedError: types.ErrInvalidDay,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			day, err := types.NewDay(tc.inputValue)

			if tc.expectError {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
				assert.Equal(t, types.Day(0), day)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedDay, day)
			}
		})
	}
}
