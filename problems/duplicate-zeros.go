package problems

// will get WA in leetcode as not in place
func DuplicateZerosNotInPlace(arr []int) []int {
	length := len(arr)

	for ix, zeros := 0, make([]int, 2); ix < length; ix++ {
		// to be appended at the end
		if ix == length-1 {
			arr = append(arr, 0)
			continue
		}

		// if 0 found for any index
		if arr[ix] == 0 {
			first := make([]int, len(arr[:ix]))
			last := make([]int, len(arr[ix+1:]))

			// split arr in two
			copy(first, arr[:ix])
			copy(last, arr[ix+1:])

			// generate arrays with zeros
			first = append(first, zeros...)
			arr = append(first, last...)

			// increase index for extra zero
			ix++
		}
	}
	arr = arr[:length]
	return arr
}

func DuplicateZerosInPlace(arr []int) []int {
	arrLen := len(arr)
	for idx := 0; idx < arrLen; idx++ {
		if idx == arrLen-1 {
			continue
		}
		if arr[idx] == 0 {
			for ix := arrLen - 2; ix > idx; ix-- {
				arr[ix+1] = arr[ix]
			}
			arr[idx+1] = 0
			idx++
		}
	}
	return arr
}
