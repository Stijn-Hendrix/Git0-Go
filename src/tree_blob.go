package main

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type TreeBlobFile struct {
	//Path string
	Name string
	Hash string
}

type TreeBlobDir struct {
	//Path      string
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
	//file.Path = path
	file.Hash = hash
	file.Name = name
	return file
}

func newTreeDir(name string) *TreeBlobDir {
	dir := new(TreeBlobDir)
	//dir.Path = path
	dir.Name = name
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

func SerializeObject(data interface{}, filename string) error {
	// Open file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Create a JSON encoder and write the data
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print with indentation
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	return nil
}

func DeserializeObject(data interface{}, filename string) (interface{}, error) {
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a JSON decoder and decode the data into the provided data structure
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return data, nil
}

// DeserializeTreeBlob reads JSON data from a file and converts it to a TreeBlobDir structure.
func DeserializeTreeBlob(filename string) *TreeBlobDir {
	var treeBlobDir TreeBlobDir
	head, err := DeserializeObject(&treeBlobDir, filename)
	if err != nil {
		log.Fatalf("Error deserializing TreeBlobDir: %v", err)
	}

	// Cast head to *TreeBlobDir
	return head.(*TreeBlobDir)
}

func DeserializeCommit(filename string) *Commit {
	var commit Commit
	head, err := DeserializeObject(&commit, filename)
	if err != nil {
		log.Fatalf("Error deserializing Commit: %v", err)
	}

	// Cast head to *Commit
	return head.(*Commit)
}
