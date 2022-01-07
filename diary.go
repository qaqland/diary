package main

import (
	"fmt"

	git "github.com/libgit2/git2go/v31"
)

func main() {
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
	date := last_commit.Author().When
	fmt.Println(date, message)

	count := last_commit.ParentCount()
	var i uint
	for i = 0; i < count; i++ {
		parent_commit := last_commit.Parent(i)
		message := parent_commit.Message()
		date := parent_commit.Author().When
		fmt.Println(date, message)
	}
}