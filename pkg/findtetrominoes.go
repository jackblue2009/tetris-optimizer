package pkg

// Here will check txt file and iterate through bytes to find tetrominoes pieces

func FindTetrominoes(convertedContent [][]byte) []Tetrominoes {
	var tetrominoToAppend Tetrominoes
	var tetrominoesList []Tetrominoes
	tetrisCounter := 1

	for a := 0; a < len(convertedContent); a++ {
		for b := 0; b < len(convertedContent[a]); b++ {
			var found bool = false
			for i := 0; i < len(PIECES); i++ {
				var skip bool = false
				for k := 0; k < len(PIECES[i]); k++ {
					for m := 0; m < len(PIECES[i][k]); m++ {
						if a+k == tetrisCounter*5-1 || b+m >= 4 {
							skip = true
							break
						}
						if convertedContent[a+k][b+m] == PIECES[i][k][m] {
							continue
						}
						skip = true
						break
					}
					if skip {
						break
					}
				}
				if !skip {
					found = true
					tetrominoToAppend.form = PIECES[i]
					tetrominoesList = append(tetrominoesList, tetrominoToAppend)
					break
				}
			}
			if found {
				if a+(tetrisCounter*5-a) >= len(convertedContent) {
					a = len(convertedContent) - 1
				} else {
					a = a + (tetrisCounter*5 - a) - 1
				}
				tetrisCounter++
				break
			}

			if a == (tetrisCounter*5)-2 && b == 3 {
				ErrorMsg()
			}
		}
	}

	return tetrominoesList
}
