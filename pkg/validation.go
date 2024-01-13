package pkg

import (
	"regexp"
	"strings"
)

func IsValidTetrisBlock(block TetrisBlock) bool {
	if len(block.Lines) != 4 {
		return false
	}

	countHash := 0
	for _, b := range block.Lines {
		if !strings.Contains(b, "#") {
			countHash++
		}
	}
	if countHash == 4 {
		return false
	}

	if !IsValidTetrisPiece(block) {
		return false
	}

	lineLength := len(block.Lines[0])
	for _, line := range block.Lines {
		if len(line) != lineLength {
			return false
		}
	}

	r := regexp.MustCompile(`^[#.]+$`)
	for _, line := range block.Lines {
		if !r.MatchString(line) {
			return false
		}
	}

	return true
}

func IsValidTetrisFile(filePath string) bool {
	tetrisBlocks, err := ReadTetrisBlocks(filePath)
	if err != nil {
		FatalError("error reading tetris file")
	}

	for _, block := range tetrisBlocks {
		if !IsValidTetrisBlock(block) {
			return false
		}
	}

	// for i := 1; i < len(tetrisBlocks); i++ {
	// 	// Only check the last line of the previous block if it is not empty
	// 	if tetrisBlocks[i-1].Lines[len(tetrisBlocks[i-1].Lines)-1] != "" && tetrisBlocks[i].Lines[0] != "" {
	// 		fmt.Println("Last line is not empty")
	// 		return false
	// 	}
	// }

	return true
}

func IsValidTetrisPiece(block TetrisBlock) bool {
	customPieces := []TetrisBlock{
		{Lines: []string{"...#", "...#", "...#", "...#"}}, //01
		{Lines: []string{"..#.", "..#.", "..#.", "..#."}}, //02
		{Lines: []string{".#..", ".#..", ".#..", ".#.."}}, //03
		{Lines: []string{"#...", "#...", "#...", "#..."}}, //04
		{Lines: []string{"....", "....", "....", "####"}}, //05
		{Lines: []string{"....", "....", "####", "...."}}, //06
		{Lines: []string{"....", "####", "....", "...."}}, //07
		{Lines: []string{"####", "....", "....", "...."}}, //08
		{Lines: []string{".###", "...#", "....", "...."}}, //09
		{Lines: []string{"....", ".###", "...#", "...."}}, //10
		{Lines: []string{"....", "....", ".###", "...#"}}, //11
		{Lines: []string{"###.", "..#.", "....", "...."}}, //12
		{Lines: []string{"....", "###.", "..#.", "...."}}, //13
		{Lines: []string{"....", "....", "###.", "..#."}}, //14
		{Lines: []string{".#..", ".#..", "##..", "...."}}, //15
		{Lines: []string{"..#.", "..#.", ".##.", "...."}}, //16
		{Lines: []string{"...#", "...#", "..##", "...."}}, //17
		{Lines: []string{"....", ".#..", ".#..", "##.."}}, //18
		{Lines: []string{"....", "..#.", "..#.", ".##."}}, //19
		{Lines: []string{"....", "...#", "...#", "..##"}}, //20
		{Lines: []string{"#...", "###.", "....", "...."}}, //21
		{Lines: []string{".#..", ".###", "....", "...."}}, //22
		{Lines: []string{"....", "#...", "###.", "...."}}, //23
		{Lines: []string{"....", ".#..", ".###", "...."}}, //24
		{Lines: []string{"....", "....", "#...", "###."}}, //25
		{Lines: []string{"....", "....", ".#..", ".###"}}, //26
		{Lines: []string{"##..", "#...", "#...", "...."}}, //27
		{Lines: []string{".##.", ".#..", ".#..", "...."}}, //28
		{Lines: []string{"..##", "..#.", "..#.", "...."}}, //29
		{Lines: []string{"....", "##..", "#...", "#..."}}, //30
		{Lines: []string{"....", ".##.", ".#..", ".#.."}}, //31
		{Lines: []string{"....", "..##", "..#.", "..#."}}, //32
		{Lines: []string{"..##", ".##.", "....", "...."}}, //32
		{Lines: []string{".##.", "##..", "....", "...."}}, //33
		{Lines: []string{"....", "..##", ".##.", "...."}}, //34
		{Lines: []string{"....", ".##.", "##..", "...."}}, //35
		{Lines: []string{"....", "....", "..##", ".##."}}, //36
		{Lines: []string{"....", "....", ".##.", "##.."}}, //37
		{Lines: []string{"#...", "##..", ".#..", "...."}}, //38
		{Lines: []string{".#..", ".##.", "..#.", "...."}}, //39
		{Lines: []string{"..#.", "..##", "...#", "...."}}, //40
		{Lines: []string{"....", "#...", "##..", ".#.."}}, //41
		{Lines: []string{"....", ".#..", ".##.", "..#."}}, //42
		{Lines: []string{"....", "..#.", "..##", "...#"}}, //43
		{Lines: []string{"##..", "##..", "....", "...."}}, //44
		{Lines: []string{".##.", ".##.", "....", "...."}}, //45
		{Lines: []string{"..##", "..##", "....", "...."}}, //46
		{Lines: []string{"....", "##..", "##..", "...."}}, //47
		{Lines: []string{"....", ".##.", ".##.", "...."}}, //48
		{Lines: []string{"....", "..##", "..##", "...."}}, //49
		{Lines: []string{"....", "....", "##..", "##.."}}, //50
		{Lines: []string{"....", "....", ".##.", ".##."}}, //51
		{Lines: []string{"....", "....", "..##", "..##"}}, //52
		{Lines: []string{"##..", ".##.", "....", "...."}}, //53
		{Lines: []string{".##.", "..##", "....", "...."}}, //54
		{Lines: []string{"....", "##..", ".##.", "...."}}, //55
		{Lines: []string{"....", ".##.", "..##", "...."}}, //56
		{Lines: []string{"....", "....", "##..", ".##."}}, //57
		{Lines: []string{"....", "....", ".##.", "..##"}}, //58
		{Lines: []string{".#..", "##..", "#...", "...."}}, //59
		{Lines: []string{"..#.", ".##.", ".#..", "...."}}, //60
		{Lines: []string{"...#", "..##", "..#.", "...."}}, //61
		{Lines: []string{"....", ".#..", "##..", "#..."}}, //62
		{Lines: []string{"....", "..#.", ".##.", ".#.."}}, //63
		{Lines: []string{"....", "...#", "..##", "..#."}}, //64
		{Lines: []string{"##..", ".#..", ".#..", "...."}}, //65
		{Lines: []string{".##.", "..#.", "..#.", "...."}}, //66
		{Lines: []string{"..##", "...#", "...#", "...."}}, //67
		{Lines: []string{"....", "##..", ".#..", ".#.."}}, //68
		{Lines: []string{"....", ".##.", "..#.", "..#."}}, //69
		{Lines: []string{"....", "..##", "...#", "...#"}}, //70
		{Lines: []string{"..#.", "###.", "....", "...."}}, //71
		{Lines: []string{"...#", ".###", "....", "...."}}, //72
		{Lines: []string{"....", "..#.", "###.", "...."}}, //73
		{Lines: []string{"....", "...#", ".###", "...."}}, //74
		{Lines: []string{"....", "....", "..#.", "###."}}, //75
		{Lines: []string{"....", "....", "...#", ".###"}}, //76
		{Lines: []string{"#...", "#...", "##..", "...."}}, //77
		{Lines: []string{".#..", ".#..", ".##.", "...."}}, //78
		{Lines: []string{"..#.", "..#.", "..##", "...."}}, //79
		{Lines: []string{"....", "#...", "#...", "##.."}}, //80
		{Lines: []string{"....", ".#..", ".#..", ".##."}}, //81
		{Lines: []string{"....", "..#.", "..#.", "..##"}}, //82
		{Lines: []string{"###.", "#...", "....", "...."}}, //83
		{Lines: []string{".###", ".#..", "....", "...."}}, //84
		{Lines: []string{"....", "###.", "#...", "...."}}, //85
		{Lines: []string{"....", ".###", ".#..", "...."}}, //86
		{Lines: []string{"....", "....", "###.", "#..."}}, //87
		{Lines: []string{"....", "....", ".###", ".#.."}}, //88
		{Lines: []string{"###.", ".#..", "....", "...."}}, //89
		{Lines: []string{".###", "..#.", "....", "...."}}, //90
		{Lines: []string{"....", "###.", ".#..", "...."}}, //91
		{Lines: []string{"....", ".###", "..#.", "...."}}, //92
		{Lines: []string{"....", "....", "###.", ".#.."}}, //93
		{Lines: []string{"....", "....", ".###", "..#."}}, //94
		{Lines: []string{".#..", "##..", ".#..", "...."}}, //95
		{Lines: []string{"..#.", ".##.", "..#.", "...."}}, //96
		{Lines: []string{"...#", "..##", "...#", "...."}}, //97
		{Lines: []string{".#..", "###.", "....", "...."}}, //98
		{Lines: []string{"..#.", ".###", "....", "...."}}, //99
		{Lines: []string{"....", ".#..", "###.", "...."}}, //100
		{Lines: []string{"....", "..#.", ".###", "...."}}, //101
		{Lines: []string{"....", "....", ".#..", "###."}}, //102
		{Lines: []string{"....", "....", "..#.", ".###"}}, //103
		{Lines: []string{"#...", "##..", "#...", "...."}}, //104
		{Lines: []string{".#..", ".##.", ".#..", "...."}}, //105
		{Lines: []string{"..#.", "..##", "..#.", "...."}}, //106
		{Lines: []string{"....", "#...", "##..", "#..."}}, //107
		{Lines: []string{"....", ".#..", ".##.", ".#.."}}, //108
		{Lines: []string{"....", "..#.", "..##", "..#."}}, //109
		{Lines: []string{".#..", "##..", "#...", "...."}}, //110
		{Lines: []string{"..#.", ".##.", ".#..", "...."}}, //111
		{Lines: []string{"...#", "..##", "..#.", "...."}}, //112
		{Lines: []string{"....", ".#..", "##..", "#..."}}, //113
		{Lines: []string{"....", "..#.", ".##.", ".#.."}}, //114
		{Lines: []string{"....", "...#", "..##", "..#."}}, //115
		{Lines: []string{"#...", "##..", ".#..", "...."}}, //116
		{Lines: []string{".#..", ".##.", "..#.", "...."}}, //117
		{Lines: []string{"..#.", "..##", "...#", "...."}}, //118
		{Lines: []string{"....", "#...", "##..", ".#.."}}, //119
		{Lines: []string{"....", ".#..", ".##.", "..#."}}, //120
		{Lines: []string{"....", "..#.", "..##", "...#"}}, //121
	}

	for _, customPiece := range customPieces {
		if IsBlockEqual(block, customPiece) {
			return true
		}
	}
	return false
}

func IsBlockEqual(block1, block2 TetrisBlock) bool {
	if len(block1.Lines) != len(block2.Lines) {
		return false
	}

	for i := 0; i < len(block1.Lines); i++ {
		if block1.Lines[i] != block2.Lines[i] {
			return false
		}
	}

	return true
}
