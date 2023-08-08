package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactorialCalculator(t *testing.T) {
	testCases := []struct {
		name               string
		a                  int
		b                  int
		expectedFactorialA uint64
		expectedFactorialB uint64
	}{
		{
			name:               "Case #1: Input a = 0 & b = 1",
			a:                  0,
			b:                  1,
			expectedFactorialA: 1,
			expectedFactorialB: 1,
		},
		{
			name:               "Case #2: Input a = 5 & b = 7",
			a:                  5,
			b:                  7,
			expectedFactorialA: 120,
			expectedFactorialB: 5040,
		},
		{
			name:               "Case #3: Input a = 11 & b = 13",
			a:                  11,
			b:                  13,
			expectedFactorialA: 39916800,
			expectedFactorialB: 6227020800,
		},
	}

	calculator := NewFactorialCalculator()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualFactorialA, actualFactorialB := calculator.Calculate(tc.a, tc.b)
			require.Equal(t, tc.expectedFactorialA, actualFactorialA)
			require.Equal(t, tc.expectedFactorialB, actualFactorialB)
		})
	}
}
