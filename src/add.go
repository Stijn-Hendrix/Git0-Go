package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func addGit0(arg string) {

	if arg == "." {
		addAll()
	}

	//serializeTreeBlob(head, "../../test.json")

	//	k, _ := deserializeTreeBlob("../../test.json")
	//fmt.Printf("%s\n", k.getHash())
}

func addAll() {
	head := dirToBlob(".")

	oldBlob := getBranchTreeBlob()

	addTreeBlob(head)

	cleanBlob(oldBlob)

	fmt.Printf("%s\n", head.getHash())
}

func dirToBlob(path string) *TreeBlobDir {

	items, _ := os.ReadDir(path)

	currentDir := newTreeDir(path)

	for _, item := range items {
		if item.Name() != ".git0" {
			filePath := filepath.Join(path, item.Name())
			if item.IsDir() {
				currentDir.addDir(dirToBlob(filePath))
			} else {
				currentDir.addFile(addFile(filePath))
			}
		}
	}
	return currentDir
}

func addFile(path string) *TreeBlobFile {
	data, _ := os.ReadFile(path)
	dataStr := string(data)
	hashStr := hashString(dataStr)

	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	if createIfNotExistsFile(".git0/objects/"+hashStr[:2]+"/"+hashStr, dataStr) {
		fmt.Printf("%s added %s\n", path, hashStr)
	}

	return newTreeFile(path, hashStr)
}

func addTreeBlob(blob *TreeBlobDir) {
	hashStr := blob.getHash()
	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	CompressAndSerialize(blob, ".git0/objects/"+hashStr[:2]+"/"+hashStr)
}
