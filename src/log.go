package main

import "fmt"

func logGit0() {
	var latestCommit *Commit

	if isHeadDetached() {
		latestCommit = getCommitFromFile(readFile(HEAD))
	} else {
		latestCommit = getCommitFromFile(getBranchLastCommitHash())
	}
	branchName := latestCommit.Branch

	const orange string = "\033[38;5;214m"
	const white string = "\033[0m"

	// Go through linked list of commits
	for {
		if latestCommit == nil {
			break
		}
		fmt.Printf(orange+"commit %s (%s)\n", latestCommit.Hash+white, branchName)
		fmt.Println()
		fmt.Printf("       %s\n", latestCommit.Message)
		fmt.Println()

		if latestCommit.Previous != "" {
			latestCommit = getCommitFromFile(latestCommit.Previous)
		} else {
			break
		}
	}
}
