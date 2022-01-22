package problems

import "strconv"

// brute force approach
func TwoSumBrute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// solution using hashmap
func TwoSumHash(nums []int, target int) []int {
	hashMap := make(map[int]string)

	for idx, val := range nums {
		k := target - val

		if hashMap[k] != "" {
			indexNew, _ := strconv.Atoi(hashMap[k])
			return []int{indexNew, idx}
		}
		if hashMap[val] == "" {
			hashMap[val] = strconv.Itoa(idx)
		}

	}
	return []int{}
}

// another solution using hashmap
func TwoSumHashAnother(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for idx, val := range nums {
		if ix, ok := hashMap[target-val]; ok {
			return []int{idx, ix}
		}
		hashMap[val] = idx

	}

	return []int{}
}
