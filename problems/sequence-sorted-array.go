package problems

import "sort"

func SortedSquares(nums []int) []int {

	sorted := make([]int, len(nums))
	for idx := range sorted {
		sorted[idx] = nums[idx] * nums[idx]
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}
