package main

import (
	"fmt"
	"os"
)

func initGit0() {
	os.RemoveAll(GIT0)

	createFolder(GIT0)
	createFolder(OBJECTS)
	createFolder(REFS)
	createFolder(REFS_HEADS)
	createFile(HEAD)
	writeBranch(MAIN_BRANCH)
	createFile(REFS_HEADS + MAIN_BRANCH)

	SerializeObject(newTreeDir("."), INDEX)

	wd, _ := os.Getwd()
	fmt.Printf("Initialized empty Git0 repository in %s\\.git0\\\n", wd)
}
