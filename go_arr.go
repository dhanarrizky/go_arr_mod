package go_arr_mod

func Reverse_char(word string) string {
	result := ""
	for i := (len(word) - 1); i >= 0; i-- {
		result = result + string(word[i])
	}
	return result
}

func Min_arr(arr []int) int {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func Max_arr(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

