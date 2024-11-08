package main

import (
	"fmt"
	"os"
)

func commitGit0() {
	blob, _ := DecompressAndDeserialize(".git0/index")
	createFiles(blob)

	hashStr := blob.getHash()
	fmt.Printf("%s", hashStr)
	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	CompressAndSerialize(blob, ".git0/objects/"+hashStr[:2]+"/"+hashStr)
}

func createFiles(t *TreeBlobDir) {

	for _, dir := range t.TreeDirs {
		createFiles(dir)
	}

	for _, file := range t.TreeFiles {
		addFile(file.Path)
	}
}

func addFile(path string) {
	data, _ := os.ReadFile(path)
	dataStr := string(data)
	hashStr := hashString(dataStr)

	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	if createIfNotExistsFile(".git0/objects/"+hashStr[:2]+"/"+hashStr, dataStr) {
		fmt.Printf("%s added %s\n", path, hashStr)
	}
}
