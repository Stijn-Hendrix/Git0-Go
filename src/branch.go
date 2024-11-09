package main

import (
	"os"
	"path/filepath"
)

func getBranchTreeBlob() *TreeBlobDir {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	treeBlob, _ := DecompressAndDeserialize(string(branchPath))
	return treeBlob
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

func getBranchTreeName() string {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	return filepath.Base(string(branchPath))
}
