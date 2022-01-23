package problems

import "fmt"

func RemoveDuplicatesTwoPointer(arr []int) (result int) {
	for i, j := 1, 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[j] = arr[i]
			j++
			result = j
		}
	}
	fmt.Println(arr)
	return
}
