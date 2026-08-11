func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	for i := range output {
		output[i] = 1
	}

	for i, n := range nums {
		if n == 1 { continue }
		for j := range output {
			if i != j {
				output[j] = output[j] * n
			}
		}
	}

	return output
}
