package main

import (
	"fmt"
	"os"
)

const MAIN_BRANCH string = "master"

func initGit0() {
	fmt.Print("Initializing git0...\n")

	os.RemoveAll(".git0")

	createFolder(".git0")
	createFolder(".git0/objects")
	createFolder(".git0/refs")
	createFolder(".git0/refs/heads")
	createFile(".git0/index", "")
	createFile(".git0/HEAD", "refs/heads/"+MAIN_BRANCH)

	fmt.Print("Git0 initialized.\n")
}
