package problems

import "fmt"

func SortArrayByParity(nums []int) []int {
	fmt.Println(nums)
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return nums
}
