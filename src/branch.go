package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createBranch(name string) bool {
	newBranchPath := ".git0/refs/heads/" + name
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
	items, _ := os.ReadDir(".git0/refs/heads/")
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

func getBranchRefsPath() string {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	return "./.git0/" + string(branchPath)
}

func getBranchLastCommitHash() string {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	refCommit, _ := os.ReadFile("./.git0/" + string(branchPath))
	return string(refCommit)
}

func getCurrentBranchName() string {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	branchName := filepath.Base(string(branchPath))
	return branchName
}

func writeBranch(name string) {
	writeToFile(".git0/HEAD", "refs/heads/"+name)
}
