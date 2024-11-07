package main

import (
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "init" {
		initGit0()
	} else if len(os.Args) == 3 && os.Args[1] == "add" && os.Args[2] == "." {
		addGit0()
	}
}
