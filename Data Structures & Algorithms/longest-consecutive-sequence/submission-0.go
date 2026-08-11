func longestConsecutive(nums []int) int {
	numMap := make(map[int]struct{})
	for _, n := range nums {
		numMap[n] = struct{}{}
	}

	runStarts := []int{}
	for _, n := range nums {
		if _, prevOk := numMap[n-1]; !prevOk {
			runStarts = append(runStarts, n)
		}
	}

	longest := 0
	for _, n := range runStarts {
		next := n+1
		current := 1
		for {
			_, nextOk := numMap[next]
			if nextOk {
				current++

				next = next+1
			} else {
				if current > longest {
					longest = current
				}
				break
			}
		}
	}
	return longest
}
