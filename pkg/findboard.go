package pkg

import "math"

// here calculate and find the minimum size possible for board

func FindBoardMinSize(tetrominoesList []Tetrominoes) int {
	minSideSize := 0
	for i := 0; i < len(tetrominoesList); i++ {
		for j := 0; j < len(tetrominoesList[i].form); j++ {
			if minSideSize < len(tetrominoesList[i].form[j]) {
				minSideSize = len(tetrominoesList[i].form[j])
			}

			if minSideSize < len(tetrominoesList[i].form) {
				minSideSize = len(tetrominoesList[i].form)
			}

			if minSideSize == 4 {
				break
			}
		}
	}

	blockCounter := math.Sqrt(float64(len(tetrominoesList) * 4))
	_, frac := math.Modf(blockCounter)
	if frac != 0 {
		blockCounter = math.Floor(blockCounter) + 1
	}

	return int(math.Max(blockCounter, float64(minSideSize)))
}
