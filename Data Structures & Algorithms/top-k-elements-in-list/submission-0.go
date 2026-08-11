func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, n := range nums {
		freq[n]++
	}

	freqList := make([][]int, len(freq))
	idx := 0
	for k,v := range freq {
		freqList[idx] = []int{k,v}
		idx++
	}

	for i := 0; i < len(freqList); i++ {
		for j := i+1; j < len(freqList); j++ {
			if freqList[i][1] < freqList[j][1] {
				freqList[i], freqList[j] = freqList[j], freqList[i]
			}
		}
	}

	largest := make([]int, k)
	for i := range k {
		largest[i] = freqList[i][0]
	}
	return largest
}
