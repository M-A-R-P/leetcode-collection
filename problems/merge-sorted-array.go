package problems

import (
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	copy(nums1[m:], nums2)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	return nums1
}
