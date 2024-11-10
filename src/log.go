package main

import "fmt"

func logGit0() {
	latestCommit := getCommitFromFile(getBranchLastCommitHash())

	const orange string = "\033[38;5;214m"
	const white string = "\033[0m"

	// Go through linked list of commits
	for {
		if latestCommit == nil {
			break
		}
		fmt.Printf(orange+"commit %s\n", latestCommit.Hash+white)
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
