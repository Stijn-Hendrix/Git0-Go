package main

import "fmt"

func checkoutGit0(arg string) {
	if commitExists(arg) {
		checkoutCommit(arg)
	} else {
		checkoutBranch(arg)
	}
}

func checkoutCommit(hash string) {

}

func checkoutBranch(name string) {
	writeBranch(name)
	fmt.Printf("Switched to branch '%s'\n", name)
}
