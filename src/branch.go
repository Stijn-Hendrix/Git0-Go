package main

import (
	"fmt"
	"os"
)

func createBranch(name string) bool {
	newBranchPath := REFS_HEADS + name
	if fileExists(newBranchPath) {
		return false
	}
	createFile(newBranchPath)
	return true
}

func branchGit0(name string) {
	if name == "" {
		logBranchStatus()
		return
	}

	if !createBranch(name) {
		fmt.Printf("Branch with name %s already exists!\n", name)
		return
	}

	logBranchStatus()
}

func logBranchStatus() {
	items, _ := os.ReadDir(REFS_HEADS)
	currentBranch := getCurrentBranchName()
	for _, item := range items {
		if !item.IsDir() {
			if item.Name() == currentBranch {
				fmt.Printf("* %s\n", item.Name())
			} else {
				fmt.Printf("  %s\n", item.Name())
			}
		}
	}
}

func writeBranch(name string) {
	writeToFile(HEAD, "refs/heads/"+name)
}

func isHeadDetached() bool {
	branchPath := readFile(HEAD)
	return commitExists(string(branchPath))
}
