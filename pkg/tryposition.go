package pkg

import "os"

// here will attempt to check position and try to place tetromino pieces per letter, and clear position

func TryPosition(piece int, tetrominoesList []Tetrominoes, size int) {
	for y := 0; y < len(BOARD); y++ {
		for x := 0; x < len(BOARD); x++ {
			if CheckPosition(y, x, piece, tetrominoesList) {
				if y == len(BOARD)-1 || piece == len(tetrominoesList)-1 {
					PrintBoard()
					os.Exit(0)
				} else {
					TryPosition(piece+1, tetrominoesList, size)
				}
				ClearPosition(y, x, piece, tetrominoesList)
			}
		}
	}

	if piece == 0 {
		increaseSize := size + 1
		CreateBoard(increaseSize)
		TryPosition(0, tetrominoesList, increaseSize)
	}
}

// func TryPosition(piece int, tetrominoesList []Tetrominoes, size int) bool {
// 	if piece == len(tetrominoesList) {
// 		PrintBoard()
// 		os.Exit(0)
// 		return true
// 	}

// 	for y := 0; y < size; y++ {
// 		for x := 0; x < size; x++ {
// 			if CheckPosition(y, x, piece, tetrominoesList) {
// 				// Place the piece and mark it as used
// 				PlacePiece(y, x, piece, tetrominoesList)

// 				// Continue with the next piece
// 				if TryPosition(piece+1, tetrominoesList, size) {
// 					return true // Solution found, exit early
// 				}

// 				// Clear the piece and mark it as unused (backtrack)
// 				ClearPosition(y, x, piece, tetrominoesList)
// 			}
// 		}
// 	}

// 	return false
// }
