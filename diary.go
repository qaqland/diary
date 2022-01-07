package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
	git "github.com/libgit2/git2go/v31"
)

var had_show mapset.Set

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
	had_show = mapset.NewSet()
	parent(last_commit)
}

func show(commit *git.Commit) {
	message := commit.Message()
	date := commit.Author().When
	fmt.Println(date, message)
}

func parent(commit *git.Commit) {
	// 如果已经 show 过，就直接返回
	if had_show.Contains(commit) {
		return
	}
	had_show.Add(commit)
	show(commit)
	count := commit.ParentCount()
	if count == 0 {
		return
	}
	var i uint
	// 可能会很多 commit 产生一个 commit
	// count = 0 走到了终点
	// count = 1 有一个爸爸
	// count = 2 有两个爸爸
	for i = 0; i < count; i++ {
		next_commit := commit.Parent(i)
		parent(next_commit)
	}
}

// go build -tags static,system_libgit2 diary.go