package binarysearch

import (
	"sort"
)

func Search(a []int, b int) bool {
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

func EasySearch(a []int, b int) bool {
	i := sort.Search(len(a), func(i int) bool { return a[i] >= b })
	if i < len(a) && a[i] == b {
		return true
	}
	return false
}

func SuperEasySearch(a []int, b int) bool {
	i := sort.SearchInts(a, b)
	if i < len(a) && a[i] == b {
		return true
	}
	return false
}
