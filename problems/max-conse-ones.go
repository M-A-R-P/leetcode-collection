package problems

func FindMaxConsecutiveOnes(nums []int) (max int) {
	localMax := max
	for _, num := range nums {
		if num != 1 {
			if localMax > max {
				max = localMax
			}
			localMax = 0
			continue
		}
		localMax++
	}
	if localMax > max {
		max = localMax
	}

	return
}
