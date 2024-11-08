package main

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
)

type TreeBlobFile struct {
	Path string
	Hash string
}

type TreeBlobDir struct {
	Path      string
	TreeDirs  []*TreeBlobDir
	TreeFiles []*TreeBlobFile
}

func (t *TreeBlobDir) getHash() string {
	hasher := sha256.New()

	hasher.Write([]byte(t.Path))

	// Sort files by path to ensure consistent hashing order
	sort.Slice(t.TreeFiles, func(i, j int) bool {
		return t.TreeFiles[i].Path < t.TreeFiles[j].Path
	})

	for _, file := range t.TreeFiles {
		hasher.Write([]byte(file.Path))
		hasher.Write([]byte(file.Hash))
	}

	// Sort subdirectories by path to ensure consistent hashing order
	sort.Slice(t.TreeDirs, func(i, j int) bool {
		return t.TreeDirs[i].Path < t.TreeDirs[j].Path
	})

	for _, subDir := range t.TreeDirs {
		hasher.Write([]byte(subDir.Path))
		hasher.Write([]byte(subDir.getHash()))
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func newTreeFile(path string, hash string) *TreeBlobFile {
	file := new(TreeBlobFile)
	file.Path = path
	file.Hash = hash
	return file
}

func newTreeDir(path string) *TreeBlobDir {
	dir := new(TreeBlobDir)
	dir.Path = path
	return dir
}

func (t *TreeBlobDir) addDir(dir *TreeBlobDir) {
	t.TreeDirs = append(t.TreeDirs, dir)
}

func (t *TreeBlobDir) addFile(file *TreeBlobFile) {
	t.TreeFiles = append(t.TreeFiles, file)
}

func CompressAndSerialize(root *TreeBlobDir, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Create a gzip writer for compression
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Create a gob encoder and write the TreeBlobDir to the gzip writer
	encoder := gob.NewEncoder(gzipWriter)
	if err := encoder.Encode(root); err != nil {
		return fmt.Errorf("failed to encode gob: %w", err)
	}

	return nil
}

func DecompressAndDeserialize(filename string) (*TreeBlobDir, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a gzip reader for decompression
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	// Create a gob decoder and decode the data into TreeBlobDir
	var root TreeBlobDir
	decoder := gob.NewDecoder(gzipReader)
	if err := decoder.Decode(&root); err != nil {
		return nil, fmt.Errorf("failed to decode gob: %w", err)
	}

	return &root, nil
}

func cleanBlob(blob *TreeBlobDir) {

}
