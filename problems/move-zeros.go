package problems

func MoveZeroes(nums []int) []int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i := len(nums) - 1; i >= count; i-- {
		nums[i] = 0
	}
	return nums
}
