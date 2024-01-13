package pkg

// Create the board that will place tetrominoes pieces

func CreateBoard(size int) {
	BOARD = nil

	for i := 0; i < size; i++ {
		BOARD = append(BOARD, nil)
		for k := 0; k < size; k++ {
			BOARD[i] = append(BOARD[i], 0)
		}
	}
}
