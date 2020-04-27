package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		arr      []int
		s        int
		expected bool
	}{
		{[]int{1, 3, 5, 7, 9, 11}, 1, true},
		{[]int{1, 3, 5, 7, 9, 11}, 2, false},
	}
	for _, tc := range testCases {
		t.Run("Basic", func(t *testing.T) {
			actual := Search(tc.arr, tc.s)

			assertSortIsCorrect(t, actual, tc.expected, tc.s, tc.arr)
		})
		t.Run("Easy", func(t *testing.T) {
			actual := EasySearch(tc.arr, tc.s)

			assertSortIsCorrect(t, actual, tc.expected, tc.s, tc.arr)
		})
		t.Run("Super easy", func(t *testing.T) {
			actual := SuperEasySearch(tc.arr, tc.s)

			assertSortIsCorrect(t, actual, tc.expected, tc.s, tc.arr)
		})
	}
}

func assertSortIsCorrect(t *testing.T, actual bool, expected bool, s int, arr []int) {
	if actual != expected {
		t.Fatalf("Search error: expected: %v, but actual: %v\nsearch %v in arr = %v\n", expected, actual, s, arr)
	}
}
