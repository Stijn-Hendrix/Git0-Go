package main

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

type TreeBlobFile struct {
	Name string
	Hash string
}

type TreeBlobDir struct {
	Name      string
	TreeDirs  []*TreeBlobDir
	TreeFiles []*TreeBlobFile
}

func (t *TreeBlobDir) getHash() string {
	hasher := sha256.New()

	hasher.Write([]byte(t.Name))

	// Sort files by path to ensure consistent hashing order
	sort.Slice(t.TreeFiles, func(i, j int) bool {
		return t.TreeFiles[i].Name < t.TreeFiles[j].Name
	})

	for _, file := range t.TreeFiles {
		hasher.Write([]byte(file.Name))
		hasher.Write([]byte(file.Hash))
	}

	// Sort subdirectories by path to ensure consistent hashing order
	sort.Slice(t.TreeDirs, func(i, j int) bool {
		return t.TreeDirs[i].Name < t.TreeDirs[j].Name
	})

	for _, subDir := range t.TreeDirs {
		hasher.Write([]byte(subDir.Name))
		hasher.Write([]byte(subDir.getHash()))
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func newTreeFile(name string, hash string) *TreeBlobFile {
	file := new(TreeBlobFile)
	file.Hash = hash
	file.Name = name
	return file
}

func newTreeDir(name string) *TreeBlobDir {
	dir := new(TreeBlobDir)
	dir.Name = name
	return dir
}

func (t *TreeBlobDir) addDir(dir *TreeBlobDir) {
	t.TreeDirs = append(t.TreeDirs, dir)
}

func (t *TreeBlobDir) addFile(file *TreeBlobFile) {
	t.TreeFiles = append(t.TreeFiles, file)
}

func getTreeFromFile(hash string) *TreeBlobDir {
	return DeserializeTreeBlob(".git0/objects/" + hash[:2] + "/" + hash)
}
