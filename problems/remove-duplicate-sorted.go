package problems

// will get WA in leetcode as not in place
func RemoveDuplicatesInPlace(nums []int) int {
	numsLen := len(nums)
	numsMap := make(map[int]int, numsLen)

	for _, val := range nums {
		if numsMap[val] == 0 {
			numsMap[val]++
		}
	}

	return len(numsMap)
}

func RemoveDuplicates(nums []int) int {
	deleteCount := 0
	numsLen := len(nums)

	for idx := 0; idx < numsLen; idx++ {
		for ix := idx + 1; ix < numsLen; {
			if nums[idx] == nums[ix] {
				if nums[idx] == 1000 {
					break
				}
				deleteCount++
				deleteAtIndex(nums, idx)
			} else {
				ix++
			}
		}
	}
	return numsLen - deleteCount
}

func deleteAtIndex(nums []int, idx int) {
	numsLen := len(nums)

	for ix := idx; ix < numsLen; ix++ {
		if ix == numsLen-1 {
			nums[ix] = 1000
			continue
		}
		nums[ix] = nums[ix+1]
	}
}
