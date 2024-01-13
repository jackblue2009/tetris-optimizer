package pkg

import "os"

func GetCmdArg() string {
	args := os.Args[1:]
	if len(args) != 1 {
		PrintUsage()
	}
	return args[0]
}
