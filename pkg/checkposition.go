package pkg

// this function will check position before attempting to place tetrominoes pieces

func CheckPosition(y int, x int, piece int, tetrominoesList []Tetrominoes) bool {
	for i := 0; i < len(tetrominoesList[piece].form); i++ {
		if len(tetrominoesList[piece].form)+y > len(BOARD) || len(tetrominoesList[piece].form[i])+x > len(BOARD) {
			return false
		}
	}

	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if tetrominoesList[piece].form[a-y][b-x] == 46 {
				continue
			}

			if BOARD[a][b] == 0 {
				BOARD[a][b] = byte(65 + piece)
			} else {
				ClearPosition(y, x, piece, tetrominoesList)
				return false
			}
		}
	}

	return true
}
