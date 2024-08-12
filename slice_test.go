package polyfill

import (
	"fmt"
	"testing"
)

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

func toStr(n int) string {
	return fmt.Sprint(n)
}

func TestMap(t *testing.T) {
	tests := []struct {
		name      string
		inputT    []int
		fT        func(int) int
		expectedT []int
		fS        func(int) string
		expectedS []string
	}{
		{
			name:      "doubles",
			inputT:    []int{1, 2, 3},
			fT:        double,
			expectedT: []int{2, 4, 6},
			fS:        toStr,
			expectedS: []string{"1", "2", "3"},
		},
		{
			name:      "squares",
			inputT:    []int{2, 4, 6},
			fT:        square,
			expectedT: []int{4, 16, 36},
			fS:        toStr,
			expectedS: []string{"2", "4", "6"},
		},
		{
			name:      "empty array",
			inputT:    []int{},
			fT:        double,
			expectedT: []int{},
			fS:        toStr,
			expectedS: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT := Map(tt.inputT, tt.fT)
			if fmt.Sprint(gotT) != fmt.Sprint(tt.expectedT) {
				t.Errorf("Map() = %v, want %v", gotT, tt.expectedT)
			}

			gotS := Map(tt.inputT, tt.fS)
			if fmt.Sprint(gotS) != fmt.Sprint(tt.expectedS) {
				t.Errorf("Map() = %v, want %v", gotS, tt.expectedS)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type test struct {
		name     string
		input    []int
		filter   func(t int) bool
		expected []int
	}

	tests := []test{
		{
			name:     "TestEmptySlice",
			input:    []int{},
			filter:   func(t int) bool { return t%2 == 0 },
			expected: []int{},
		},
		{
			name:     "TestOddNumbers",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			filter:   func(t int) bool { return t%2 != 0 },
			expected: []int{1, 3, 5, 7},
		},
		{
			name:     "TestEvenNumbers",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			filter:   func(t int) bool { return t%2 == 0 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "TestZeroElement",
			input:    []int{0, 1, 2, 3, 4, 5},
			filter:   func(t int) bool { return t == 0 },
			expected: []int{0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Filter(tc.input, tc.filter)
			for i, v := range result {
				if tc.expected[i] != v {
					t.Errorf("%s: Expected: %v, Result: %v", tc.name, tc.expected, result)
				}
			}
		})
	}
}
