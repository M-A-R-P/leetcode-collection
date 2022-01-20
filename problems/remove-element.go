package problems

func RemoveElement(nums []int, val int) int {
	emptySpots := 0
	numsLen := len(nums)
	for idx := 0; idx < numsLen; {
		if nums[idx] == val {
			for ix := idx; ix < numsLen; ix++ {
				if ix == numsLen-1 {
					nums[ix] = 1000
					emptySpots++
					continue
				}
				nums[ix] = nums[ix+1]
			}
		} else {
			idx++
		}
	}
	return numsLen - emptySpots
}
