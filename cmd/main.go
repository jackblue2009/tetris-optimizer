package main

import (
	"bytes"
	"os"
	"tetris-optimizer/pkg"
)

func RemoveEmptySpace(list [][]byte) [][]byte {
	var newList [][]byte

	for _, l := range list {
		if len(l) > 0 {
			newList = append(newList, l)
		}
	}

	return newList
}

func AddEmptySlices(list [][]byte, interval int) [][]byte {
	var updatedSlice [][]byte

	for i, slice := range list {
		updatedSlice = append(updatedSlice, slice)

		// Check if we should add an empty slice after every 'interval' slices.
		if (i+1)%interval == 0 && i != len(list)-1 {
			updatedSlice = append(updatedSlice, []byte{})
		}
	}

	return updatedSlice
}

func main() {
	content, err := os.ReadFile(pkg.GetCmdArg())
	if err != nil {
		pkg.FatalError("error reading file...")
	}
	if !pkg.IsValidTetrisFile(pkg.GetCmdArg()) {
		pkg.ErrorMsg()
	}

	// separate content into bytes and store in new variable
	div := []byte{10}
	newContent := bytes.Split(content, div)

	list := RemoveEmptySpace(newContent)
	list = AddEmptySlices(list, 4)

	// check format of new slice
	//pkg.CheckFormat(list)

	// find valid tetrominoes pieces and calculate minimum board size
	tetrominoesList := pkg.FindTetrominoes(list)
	minSize := pkg.FindBoardMinSize(tetrominoesList)

	// create board and place pieces
	pkg.CreateBoard(minSize)
	pkg.TryPosition(0, tetrominoesList, minSize)
}
