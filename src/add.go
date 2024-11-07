package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func addGit0() {

	/*
		head, _ := os.ReadFile("./HEAD")
		headStr := string(head)

		var treeBlobHash string

		if fileExists(headStr) {
			treeBlob, _ := os.ReadFile(headStr)
			treeBlobHash = string(treeBlob)
		}

		if len(head) == 0 {
			writeToFile("./HEAD", "TEST")
		}
	*/

	dirToBlob(".")
}

func dirToBlob(path string) *TreeBlobDir {

	items, _ := os.ReadDir(path)

	currentDir := newTreeDir(path)

	for _, item := range items {
		if item.Name() != ".git0" {
			filePath := filepath.Join(path, item.Name())
			if item.IsDir() {
				fmt.Printf("Dir: %s\n", filePath)
				currentDir.addDir(dirToBlob(filePath))
			} else {
				addFile(filePath)
				fmt.Printf("File: %s\n", filePath)
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
		fmt.Printf("%s added", path)
	}

	return newTreeFile(path, dataStr, hashStr)
}
