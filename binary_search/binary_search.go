package binary_search

import "sort"

func Search(a []int, b int) bool {
	sort.Ints(a)
	front := 0
	back := len(a)
	for front <= back {
		mid := (front + back) / 2
		if a[mid] == b {
			return true
		}
		if a[mid] > b {
			back = mid - 1
		} else {
			front = mid + 1
		}
	}
	return false
}
