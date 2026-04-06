package examples

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}

	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}

	return result
}