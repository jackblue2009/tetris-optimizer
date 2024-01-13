package pkg

// this function will check format of txt file if it is valid blocks

func CheckFormat(convertedContent [][]byte) {
	index := 0

	for i := 0; i < len(convertedContent); i++ {
		if len(convertedContent[i]) == 0 {
			if index != 4 {
				ErrorMsg()
			} else {
				index = 0
			}
		}
		for k := 0; k < len(convertedContent[i]); k++ {
			if convertedContent[i][k] == 35 {
				index++
			}
		}
	}
}
