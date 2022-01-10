package main

import (
	"flag"
	"log"

	mapset "github.com/deckarep/golang-set"
	git "github.com/libgit2/git2go/v31"
)

var Link_commit []*git.Commit

var Pagination int
var Public string

var had_get mapset.Set

func main() {

	// 命令行参数
	flag.IntVar(&Pagination, "p", 20, "pagination")
	flag.StringVar(&Public, "o", "./public", "dir for export")
	flag.Parse()

	// 复制文件
	start_public()

	// 准备 set 防止便历重复
	had_get = mapset.NewSet()

	// git 操作
	git_diary, err := git.OpenRepository("./")
	if err != nil {
		log.Println("open git repo err:", err)
		return
	}

	// 头节点
	head, err := git_diary.Head()
	if err != nil {
		log.Println("get repo head err:", err)
		return
	}
	oid := head.Target()

	// 得到最后一次 commit
	last_commit, err := git_diary.LookupCommit(oid)
	if err != nil {
		log.Println("lookup commit err:", err)
		return
	}

	// 开始遍历
	get_parent(last_commit)

	// 解析
	pagination_tmpl()
}

func get_parent(commit *git.Commit) {

	// 如果已经 get 过，就直接返回
	if had_get.Contains(commit) {
		return
	}

	// 添加 set
	had_get.Add(commit)

	// 得到 这一次的 commit
	Link_commit = append(Link_commit, commit)
	
	// 一个 commit 可能由多个爸爸产生
	// count = 0 走到了终点
	for i := uint(0); i < commit.ParentCount(); i++ {
		next_commit := commit.Parent(i)
		get_parent(next_commit)
		// 这里曾经出现过 bug 不知道为啥又好了，不要动
	}
}

// go build -tags static *.go
