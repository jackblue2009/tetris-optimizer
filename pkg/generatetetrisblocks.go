package pkg

func GenerateTetrisBlocks(blocks []TetrisBlock) []TetrisBlock {
	newBlocks := []TetrisBlock{}
	letter := 'A'
	for _, block := range blocks {
		newBlock := TetrisBlock{Lines: []string{}}

		for _, line := range block.Lines {
			newLine := ""
			for _, char := range line {
				if char == '#' {
					newLine += string(letter)
				} else {
					newLine += string(char)
				}
			}

			newBlock.Lines = append(newBlock.Lines, newLine)
		}

		newBlocks = append(newBlocks, newBlock)
		letter++
	}

	return newBlocks
}
