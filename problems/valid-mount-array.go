package problems

import "sort"

func ValidMountainArray(arr []int) bool {
	arrLen := len(arr)

	// length less than 3, can't be mountain
	if arrLen < 3 {
		return false
	}

	// if sorted, can't be mountain
	ascSorted := sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	descSorted := sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	if ascSorted || descSorted {
		return false
	}

	for i, j := 0, arrLen-1; i < arrLen-1 && j >= 0; {
		// break if peak point found
		if i == j && arr[i] == arr[j] {
			return true
		}

		// check if the array is sorted in asc/desc order
		if arr[i] < arr[i+1] {
			i++
			continue
		}
		if arr[j] < arr[j-1] {
			j--
			continue
		}

		return false
	}

	return false
}
