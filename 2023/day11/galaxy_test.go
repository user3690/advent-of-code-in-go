package day11

import "testing"

func TestCreatePairs(t *testing.T) {
	tests := []struct {
		galaxies []position
		expected int
	}{
		{
			galaxies: []position{
				{
					row: 0,
					col: 0,
				},
				{
					row: 1,
					col: 0,
				},
				{
					row: 0,
					col: 1,
				},
				{
					row: 1,
					col: 1,
				},
			},
			expected: 6,
		},
		{
			galaxies: []position{
				{
					row: 0,
					col: 0,
				},
				{
					row: 1,
					col: 0,
				},
				{
					row: 0,
					col: 1,
				},
			},
			expected: 3,
		},
		{
			galaxies: []position{
				{
					row: 0,
					col: 0,
				},
				{
					row: 1,
					col: 0,
				},
			},
			expected: 1,
		},
		{
			galaxies: []position{
				{
					row: 0,
					col: 0,
				},
				{
					row: 0,
					col: 0,
				},
			},
			expected: 0,
		},
		{
			galaxies: []position{
				{
					row: 0,
					col: 0,
				},
				{
					row: 1,
					col: 0,
				},
				{
					row: 0,
					col: 1,
				},
				{
					row: 1,
					col: 1,
				},
				{
					row: 2,
					col: 2,
				},
			},
			expected: 10,
		},
	}

	for _, test := range tests {
		pairs := createPairs(test.galaxies)
		if len(pairs) != test.expected {
			t.Fatalf("expected length %d but got %d", test.expected, len(pairs))
		}
	}
}
