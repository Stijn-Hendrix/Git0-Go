package main

import (
	"fmt"
	"os"
)

func writeCheckoutCommit(hash string) {
	writeToFile(HEAD, hash)
}

func checkoutGit0(arg string) {

	if arg == "" {
		arg = getCurrentBranchName()
	}

	if commitExists(arg) {
		checkoutCommit(arg)
	} else {
		checkoutBranch(arg)
	}
}

func checkoutCommit(hash string) {
	clearWorkingDir()

	commit := getCommitFromFile(hash)
	tree := getTreeFromFile(commit.Tree)
	buildWorkingDir(tree, ".")

	headHashOfBranch := getBranchLastCommitHashFrom(commit.Branch)
	isHead := headHashOfBranch == commit.Hash
	if !isHead {
		writeCheckoutCommit(commit.Hash)
	}
}

func checkoutBranch(name string) {
	writeBranch(name)

	latestCommit := getBranchLastCommitHash()
	if commitExists(latestCommit) {
		checkoutCommit(latestCommit)
	}

	fmt.Printf("Switched to branch '%s'\n", name)
}

func clearWorkingDir() {
	items, _ := os.ReadDir(".")

	for _, item := range items {
		if item.Name() != ".git0" {
			os.RemoveAll("./" + item.Name())
		}
	}
}

func buildWorkingDir(t *TreeBlobDir, rootPath string) {

	for _, file := range t.TreeFiles {
		currentPath := rootPath + "/" + file.Name
		createFileAndWrite(currentPath, readSavedFile(file.Hash))
	}

	for _, dir := range t.TreeDirs {
		currentPath := rootPath + "/" + dir.Name
		createFolder(currentPath)
		buildWorkingDir(dir, currentPath)
	}

}
