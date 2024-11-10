package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func commitGit0(message string) {

	// Get index tree
	indexBlob := DeserializeTreeBlob(INDEX)
	indexBlobHash := indexBlob.getHash()

	// Write new commit to objects
	newCommit := newCommit(indexBlobHash, message, getBranchLastCommitHash(), getCurrentBranchName())

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
	createIfNotExistsFolder(objectDirPath(newCommit.Hash))
	SerializeObject(newCommit, objectFilePath(newCommit.Hash))

	// Write tree blob to file
	createIfNotExistsFolder(objectDirPath(indexBlobHash))
	SerializeObject(indexBlob, objectFilePath(indexBlobHash))

	// Create files in objects from index tree
	createFiles(indexBlob, ".")

	// Write new latest commit to refs
	writeToFile(getBranchRefsPath(), newCommit.Hash)

	// Re-init index
	SerializeObject(newTreeDir("."), INDEX)
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
	dataStr := readFile(path)
	hashStr := hashString(dataStr)

	createIfNotExistsFolder(objectDirPath(hashStr))
	if createIfNotExistsFile(objectFilePath(hashStr), dataStr) {
		fmt.Printf("create node .git0/objects/%s/%s (%s)\n", hashStr[:2], hashStr, path)
	}
}
