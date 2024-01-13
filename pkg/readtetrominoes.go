package pkg

import (
	"bufio"
	"os"
	"strings"
)

func ReadTetrisBlocks(filePath string) ([]TetrisBlock, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tetrisBlocks []TetrisBlock
	var currentBlock TetrisBlock

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if len(currentBlock.Lines) > 0 {
				tetrisBlocks = append(tetrisBlocks, currentBlock)
			}
			currentBlock = TetrisBlock{}
		} else {
			currentBlock.Lines = append(currentBlock.Lines, line)
		}
	}

	if len(currentBlock.Lines) > 0 {
		tetrisBlocks = append(tetrisBlocks, currentBlock)
	}

	return tetrisBlocks, nil
}
