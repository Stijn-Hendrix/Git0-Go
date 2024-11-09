package main

import (
	"fmt"
	"os"
)

type Commit struct {
	Message  string
	Hash     string
	Previous string // Hash pointer
	Tree     *TreeBlobDir
}

func newCommit(t *TreeBlobDir, message string, previous string) *Commit {
	commit := new(Commit)
	commit.Tree = t
	commit.Message = message
	commit.Hash = t.getHash()
	commit.Previous = previous
	return commit
}

func getCommitFromFile(hash string) *Commit {
	return DeserializeCommit(".git0/objects/" + hash[:2] + "/" + hash)
}

func commitGit0(message string) {

	// Get index tree
	blob := DeserializeTreeBlob(".git0/index")
	hashStr := blob.getHash()

	fmt.Printf("[%s %s] commit\n", getBranchTreeName(), hashStr)

	// Create files in index tree
	createFiles(blob, ".")

	// Create commit folder
	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])

	// Write new commit to objects
	newCommit := newCommit(blob, message, getBranchLastCommitHash())
	SerializeObject(newCommit, ".git0/objects/"+hashStr[:2]+"/"+hashStr)

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
