func isValidSudoku(board [][]byte) bool {
	// check rows 
	for y := range board {
		seen := make(map[byte]struct{})
		for x := range board {
			v := board[y][x]	
			if v == '.' { continue }
			if _, ok := seen[v]; ok {
				return false
			}
			seen[v] = struct{}{}
		}
	}

	// check cols 
	for x := range board {
		seen := make(map[byte]struct{})
		for y := range board {
			v := board[y][x]	
			if v == '.' { continue }
			if _, ok := seen[v]; ok {
				return false
			}
			seen[v] = struct{}{}
		}
	}

	// check 3x3s
	for bx := 0; bx < 3; bx++ {
		for by := 0; by < 3; by++ {
			seen := make(map[byte]struct{})
				for x := 0; x < 3; x++ {
					for y := 0; y < 3; y++ {
						v := board[by*3+y][bx*3+x]
						if v == '.' { continue }
						if _, ok := seen[v]; ok {
							return false
						}
						seen[v] = struct{}{}
					}
				}
		}
	}

	return true
}
