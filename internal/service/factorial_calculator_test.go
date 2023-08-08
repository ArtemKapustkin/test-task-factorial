package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFactorialCalculator(t *testing.T) {
	type payload struct {
		a int
		b int
	}

	testCases := []struct {
		name               string
		payload            payload
		expectedFactorialA uint64
		expectedFactorialB uint64
	}{
		{
			name: "Case #1: Input a = 0 & b = 1",
			payload: payload{
				a: 0,
				b: 1,
			},
			expectedFactorialA: 1,
			expectedFactorialB: 1,
		},
		{
			name: "Case #2: Input a = 5 & b = 7",
			payload: payload{
				a: 5,
				b: 7,
			},
			expectedFactorialA: 120,
			expectedFactorialB: 5040,
		},
		{
			name: "Case #3: Input a = 11 & b = 13",
			payload: payload{
				a: 11,
				b: 13,
			},
			expectedFactorialA: 39916800,
			expectedFactorialB: 6227020800,
		},
	}

	calculator := NewFactorialCalculator()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualFactorialA, actualFactorialB := calculator.Calculate(tc.payload.a, tc.payload.b)

			require.Equal(t, tc.expectedFactorialA, actualFactorialA)
			require.Equal(t, tc.expectedFactorialB, actualFactorialB)
		})
	}
}
