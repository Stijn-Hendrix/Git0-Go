package main

type TreeBlobFile struct {
	Path    string
	Content string
	Hash    string
}

type TreeBlobDir struct {
	Path      string
	TreeDirs  []*TreeBlobDir
	TreeFiles []*TreeBlobFile
}

/*
func (t *TreeBlobDir) getHash() string {

}
*/

func newTreeFile(path string, content string, hash string) *TreeBlobFile {
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
