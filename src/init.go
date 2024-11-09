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
	createFile(".git0/HEAD", "refs/heads/"+MAIN_BRANCH)
	createFile(".git0/refs/heads/"+MAIN_BRANCH, "")

	SerializeObject(newTreeDir("."), ".git0/index")
	//createFile(".git0/index", "")

	fmt.Print("Git0 initialized.\n")
}
