package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func addGit0(path string) {

	head, _ := DeserializeTreeBlob(".git0/index")

	if path == "." {
		SerializeTreeBlob(createDirBlob("."), ".git0/index")
	} else {
		if path[0] != '.' {
			path = "." + path
		}

		fmt.Printf("%s\n", path)

		if isFile(path) {
			addFileToTree(path, head)
		} else {
			addDirToTree(path, head)
		}

		SerializeTreeBlob(head, ".git0/index")
	}
}

func findOrCreateDir(t *TreeBlobDir, dirName string) *TreeBlobDir {

	for _, dir := range t.TreeDirs {
		if dir.Name == dirName {
			return dir
		}
	}

	newDir := newTreeDir(dirName)
	t.addDir(newDir)
	return newDir
}

func addDirToTree(path string, t *TreeBlobDir) {
	fmt.Println(path)
	dir := createDirBlob(path)

	// Split file path into directory components
	dirs := strings.Split(filepath.Dir(path), string(os.PathSeparator))

	fmt.Println(dirs)

	currentDir := t

	// Traverse the directory structure, creating directories as needed
	for _, dirName := range dirs {
		if dirName == "" {
			continue
		}
		currentDir = findOrCreateDir(currentDir, dirName)
		fmt.Println(currentDir.Name)
	}

	addOrReplaceDir(dir, currentDir)
}

func addFileToTree(path string, t *TreeBlobDir) {
	file := createFileBlob(path)

	// Split file path into directory components
	dirs := strings.Split(filepath.Dir(path), string(os.PathSeparator))
	currentDir := t

	// Traverse the directory structure, creating directories as needed
	for _, dirName := range dirs {
		if dirName == "" {
			continue
		}
		currentDir = findOrCreateDir(currentDir, dirName)
	}

	addOrReplaceFile(currentDir, file)
}

func addOrReplaceFile(dir *TreeBlobDir, file *TreeBlobFile) {

	for _, f := range dir.TreeFiles {
		if f.Hash == file.Hash {
			return
		}
	}

	dir.TreeFiles = append(dir.TreeFiles, file)
}

func addOrReplaceDir(dir *TreeBlobDir, root *TreeBlobDir) {

	for i, d := range root.TreeDirs {
		if d.Name == dir.Name {
			root.TreeDirs[i] = dir
			return
		}
	}

	root.TreeDirs = append(root.TreeDirs, dir)
}

func createFileBlob(path string) *TreeBlobFile {
	data, _ := os.ReadFile(path)
	dataStr := string(data)
	hashStr := hashString(dataStr)
	file := newTreeFile(filepath.Base(path), hashStr)
	return file
}

func createDirBlob(path string) *TreeBlobDir {
	items, _ := os.ReadDir(path)

	currentDir := newTreeDir(filepath.Base(path))

	for _, item := range items {
		if item.Name() != ".git0" {
			filePathExtended := filepath.Join(path, item.Name())
			if item.IsDir() {
				currentDir.addDir(createDirBlob(filePathExtended))
			} else {
				currentDir.addFile(createFileBlob(filePathExtended))
			}
		}
	}
	return currentDir
}
