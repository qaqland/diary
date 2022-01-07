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
	show(last_commit)

	count := last_commit.ParentCount()
	parent_commit := last_commit.Parent(0)
	// 可能会很多 commit 产生一个 commit
	// count = 0 走到了重点
	// count = 1 有一个爸爸
	// count = 2 有两个爸爸

	var i uint
	for i = 0; i < count; i++ {
		parent_commit = parent_commit.Parent(0)
		show(parent_commit)
	}

}
func show(commit *git.Commit) {
	message := commit.Message()
	date := commit.Author().When
	fmt.Println(date, message)
}
