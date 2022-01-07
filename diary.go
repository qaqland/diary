package main

import (
	"fmt"

	git "github.com/libgit2/git2go/v31"
)

func main() {
	fmt.Println("123")
	diary, err := git.OpenRepository("./")
	if err != nil {
		fmt.Println("open git repo err:", err)
	}
	head, err := diary.Head()
	oid := head.Target()
	last_commit, err := diary.LookupCommit(oid)
	if err != nil {
		fmt.Println("lookup commit err:", err)
	}
	message := last_commit.Message()

	fmt.Println(message)
}
