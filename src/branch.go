package main

import "os"

func getBranchTreeBlob() *TreeBlobDir {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	treeBlob, _ := DecompressAndDeserialize(string(branchPath))
	return treeBlob
}

func getBranchTreeBlobHash() string {
	branchPath, _ := os.ReadFile("./.git0/HEAD")
	treeBlob, _ := os.ReadFile(string(branchPath))
	return string(treeBlob)
}
