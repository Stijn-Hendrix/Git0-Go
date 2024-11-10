package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

type Commit struct {
	Message  string
	Hash     string
	Previous string // Hash pointer
	Branch   string
	Tree     string // Hash pointer
}

func newCommit(tree string, message string, previous string, branch string) *Commit {
	commit := new(Commit)
	commit.Tree = tree
	commit.Message = message
	commit.Previous = previous
	commit.Branch = branch

	hasher := sha256.New()
	hasher.Write([]byte(tree))
	hasher.Write([]byte(branch))

	commit.Hash = hex.EncodeToString(hasher.Sum(nil))
	return commit
}

func getCommitFromFile(hash string) *Commit {
	return DeserializeCommit(".git0/objects/" + hash[:2] + "/" + hash)
}

func commitExists(hash string) bool {
	if len(hash) <= 2 {
		return false
	}
	return fileExists(".git0/objects/" + hash[:2] + "/" + hash)
}

func commitGit0(message string) {

	// Get index tree
	blob := DeserializeTreeBlob(".git0/index")
	hashStr := blob.getHash()

	// Write new commit to objects
	newCommit := newCommit(hashStr, message, getBranchLastCommitHash(), getCurrentBranchName())

	if commitExists(newCommit.Hash) {
		fmt.Println("Nothing to commit!")
		return
	}

	if isHeadDetached() {
		fmt.Printf("You are in 'Detached HEAD' state. If you want to create a new branch to retain commits, use \ngit0 branch new_branch_name\ngit0 checkout new_branch_name\n")
		return
	}

	fmt.Printf("[%s %s] commit\n", getCurrentBranchName(), newCommit.Hash)

	// Write commit to file
	createIfNotExistsFolder(".git0/objects/" + newCommit.Hash[:2])
	SerializeObject(newCommit, ".git0/objects/"+newCommit.Hash[:2]+"/"+newCommit.Hash)

	// Write tree blob to file
	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	SerializeObject(blob, ".git0/objects/"+hashStr[:2]+"/"+hashStr)

	// Create files in index tree
	createFiles(blob, ".")

	// Write new latest commit to refs
	writeToFile(getBranchRefsPath(), newCommit.Hash)

	// Re-init index
	SerializeObject(newTreeDir("."), ".git0/index")
}

func createFiles(t *TreeBlobDir, path string) {

	for _, dir := range t.TreeDirs {
		createFiles(dir, path+"\\"+dir.Name)
	}
	for _, file := range t.TreeFiles {
		addFile(path + "\\" + file.Name)
	}
}

func addFile(path string) {
	data, _ := os.ReadFile(path)
	dataStr := string(data)

	hashStr := hashString(dataStr)

	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	if createIfNotExistsFile(".git0/objects/"+hashStr[:2]+"/"+hashStr, dataStr) {
		fmt.Printf("create node .git0/objects/%s/%s (%s)\n", hashStr[:2], hashStr, path)
	}
}
