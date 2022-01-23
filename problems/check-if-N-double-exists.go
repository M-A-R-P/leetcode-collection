package problems

import (
	"fmt"
	"sort"
)

func CheckIfExist(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for idx, val := range arr {
		if binarySearch(arr, 2*val, idx, 0, len(arr)-1) {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, val int, idx int, low int, high int) bool {
	if low > high {
		return false
	}

	mid := int((low + (high-low)/2))

	if arr[mid] == val && mid != idx {
		fmt.Println(arr[mid])
		return true
	}
	if val > arr[mid] {
		return binarySearch(arr, val, idx, mid+1, high)
	}

	return binarySearch(arr, val, idx, low, mid-1)
}
