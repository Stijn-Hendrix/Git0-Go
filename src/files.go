package main

import (
	"fmt"
	"log"
	"os"
)

func createFolder(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
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
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func createIfNotExistsFile(path string, content string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return false
	}

	if os.IsNotExist(err) {
		createFile(path, content)
		return true
	} else {
		return false
	}
}

func createFile(path string, content string) {
	destination, err := os.Create(path)
	if err != nil {
		return
	}
	defer destination.Close()
	fmt.Fprintf(destination, "%s ", content)
}
