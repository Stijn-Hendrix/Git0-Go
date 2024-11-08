package main

import (
	"os"
	"path/filepath"
)

func addGit0(arg string) {

	if arg == "." {
		addAll()
	} else {

		file, _ := os.Open(arg)
		fileInfo, _ := file.Stat()
		defer file.Close()

		if fileInfo.IsDir() {

		} else {
			addFile(arg)
		}
	}
}

func addAll() {
	head := dirToBlob(".")
	CompressAndSerialize(head, ".git0/index")
}

func addFileToBlob(path string) {

}

func dirToBlob(path string) *TreeBlobDir {
	items, _ := os.ReadDir(path)

	currentDir := newTreeDir(path, filepath.Base(path))

	for _, item := range items {
		if item.Name() != ".git0" {
			filePath := filepath.Join(path, item.Name())
			if item.IsDir() {
				currentDir.addDir(dirToBlob(filePath))
			} else {
				data, _ := os.ReadFile(filePath)
				dataStr := string(data)
				hashStr := hashString(dataStr)
				currentDir.addFile(newTreeFile(filePath, item.Name(), hashStr))
			}
		}
	}
	return currentDir
}
