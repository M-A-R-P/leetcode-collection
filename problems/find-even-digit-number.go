package problems

import "strconv"

func FindNumbers(nums []int) (sum int) {
	for _, num := range nums {
		length := len(strconv.Itoa(num))
		if length%2 == 0 {
			sum++
		}
	}
	return
}
