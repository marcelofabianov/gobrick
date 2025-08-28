package types_test

import (
	"testing"
	"time"

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
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedDay, day)
			}
		})
	}
}

func TestDay_TimeCalculations(t *testing.T) {
	// Usamos uma data fixa para tornar os testes determinísticos.
	// 15 de Março de 2024
	today := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)

	t.Run("HasPassed", func(t *testing.T) {
		assert.True(t, types.Day(10).HasPassed(today))
		assert.False(t, types.Day(15).HasPassed(today))
		assert.False(t, types.Day(20).HasPassed(today))
	})

	t.Run("DaysUntil", func(t *testing.T) {
		assert.Equal(t, 0, types.Day(15).DaysUntil(today), "should be 0 days until the same day")
		assert.Equal(t, 5, types.Day(20).DaysUntil(today), "should be 5 days until a future day in the same month")
		// Março tem 31 dias. (31 - 15) + 10 = 26
		assert.Equal(t, 26, types.Day(10).DaysUntil(today), "should calculate days until a day in the next month")
	})

	t.Run("DaysOverdue", func(t *testing.T) {
		assert.Equal(t, 0, types.Day(15).DaysOverdue(today), "should be 0 days overdue for the same day")
		assert.Equal(t, 5, types.Day(10).DaysOverdue(today), "should be 5 days overdue for a past day in the same month")
		// Fevereiro de 2024 teve 29 dias. (29 - 20) + 15 = 24
		assert.Equal(t, 24, types.Day(20).DaysOverdue(today), "should calculate days overdue from the previous month")
	})
}
