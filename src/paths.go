package main

const MAIN_BRANCH string = "master"
const GIT0 string = ".git0/"
const OBJECTS string = GIT0 + "objects/"
const REFS string = GIT0 + "refs/"
const REFS_HEADS string = GIT0 + "refs/heads/"
const HEAD string = GIT0 + "HEAD"
const INDEX string = GIT0 + "index"

func objectDirPath(hash string) string {
	return OBJECTS + hash[:2]
}

func objectFilePath(hash string) string {
	return objectDirPath(hash) + "/" + hash
}
