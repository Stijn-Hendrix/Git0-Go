package main

import (
	"fmt"
	"os"
)

func commitGit0() {
	fmt.Println("Commiting staging area...")

	blob, _ := DeserializeTreeBlob(".git0/index")
	createFiles(blob, ".")

	hashStr := blob.getHash()
	createIfNotExistsFolder(".git0/objects/" + hashStr[:2])
	SerializeTreeBlob(blob, ".git0/objects/"+hashStr[:2]+"/"+hashStr)

	SerializeTreeBlob(newTreeDir("."), ".git0/index")

	fmt.Printf("Commited staging area (%s)", hashStr)
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
		fmt.Printf("%s added %s\n", path, hashStr)
	}
}
