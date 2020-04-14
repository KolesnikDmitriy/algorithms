package binary_search

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		arr      []int
		s        int
		expected bool
	}{
		{[]int{5, 8, 2, 4, 7, 2}, 1, false},
		{[]int{5, 8, 2, 4, 7, 2}, 2, true},
	}
	for _, tc := range testCases {
		actual := Search(tc.arr, tc.s)

		if actual != tc.expected {
			t.Fatalf("Search error: expected: %v, but actual: %v\nsearch %v in arr = %v\n", tc.expected, actual, tc.s, tc.arr)
		}
	}

}
