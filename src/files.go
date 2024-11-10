package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFolder(path string) {
	err := os.Mkdir(path, 0755)
	check(err)
}

func writeToFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), os.ModePerm)
}

func createIfNotExistsFolder(path string) {
	_, err := os.Stat(path)
	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		check(err)
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createIfNotExistsFile(path string, content string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return false
	}

	if os.IsNotExist(err) {
		createFileAndWrite(path, content)
		return true
	} else {
		return false
	}
}

func createFileAndWrite(path string, content string) {
	destination, err := os.Create(path)
	if err != nil {
		return
	}
	defer destination.Close()
	fmt.Fprintf(destination, "%s ", content)
}

func createFile(path string) {
	destination, err := os.Create(path)
	if err != nil {
		return
	}
	defer destination.Close()
}

func isFile(path string) bool {
	file, _ := os.Open(path)
	fileInfo, _ := file.Stat()
	defer file.Close()
	return !fileInfo.IsDir()
}

func getBranchRefsPath() string {
	branchPath, _ := os.ReadFile(HEAD)
	return GIT0 + string(branchPath)
}

func getBranchLastCommitHash() string {
	branchPath, _ := os.ReadFile(HEAD)
	refCommit, _ := os.ReadFile(GIT0 + string(branchPath))
	return string(refCommit)
}

func getBranchLastCommitHashFrom(branch string) string {
	refCommit, _ := os.ReadFile(REFS_HEADS + string(branch))
	return string(refCommit)
}

func getCurrentBranchName() string {
	branchPath := readFile(HEAD)
	branchName := filepath.Base(string(branchPath))

	if commitExists(branchName) {
		return getCommitFromFile(branchName).Branch
	}

	return branchName
}

func readSavedFile(hash string) string {
	path := objectFilePath(hash)
	return readFile(path)
}

func readFile(path string) string {
	data, _ := os.ReadFile(path)
	dataStr := string(data)
	return dataStr
}

func getCommitFromFile(hash string) *Commit {
	return DeserializeCommit(objectFilePath(hash))
}

func commitExists(hash string) bool {
	if len(hash) <= 2 {
		return false
	}
	return fileExists(objectFilePath(hash))
}

func getTreeFromFile(hash string) *TreeBlobDir {
	return DeserializeTreeBlob(objectFilePath(hash))
}
