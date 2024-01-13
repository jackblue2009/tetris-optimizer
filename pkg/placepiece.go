package pkg

func PlacePiece(y, x, piece int, tetrominoesList []Tetrominoes) {
	tetromino := tetrominoesList[piece]
	for i := 0; i < len(tetromino.form); i++ {
		for j := 0; j < len(tetromino.form[i]); j++ {
			if tetromino.form[i][j] == 35 { // '#' represents the occupied cell in the tetromino
				BOARD[y+i][x+j] = byte(65 + piece) // Mark the cell as used with the corresponding letter (A, B, C, ...)
			}
		}
	}
}
