package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	git "github.com/libgit2/git2go/v31"
)

type Fianl struct {
	Final_commit []*git.Commit
	Here         int
	Len          int
}

func pagination_tmpl() {
	len := len(Link_commit)
	I, K := len/Pagination, len%Pagination

	// 开始的几页完整的
	final_c := make([]*git.Commit, Pagination)
	for i := 0; i < I; i++ {
		for k := 0; k < Pagination; k++ {
			final_c[k] = Link_commit[i*Pagination+k]
		}
		final_s := Fianl{final_c, i + 1, I + 1}
		if err := write(final_s); err != nil {
			log.Println(err)
			return
		}
	}
	
	// 最后一页
	final_c_end := make([]*git.Commit, K)
	for k := 0; k < K; k++ {
		final_c_end[k] = Link_commit[I*Pagination+k]
	}
	final_s := Fianl{final_c_end, I + 1, I + 1} // 未测试
	if err := write(final_s); err != nil {
		log.Println(err)
		return
	}
}

func write(final Fianl) error {
	html_index := strconv.Itoa(final.Here)
	if final.Here == 1 {
		html_index = "index"
	}
	public_index, err := os.OpenFile(Public+"/"+html_index+".html",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer public_index.Close()
	index_tmpl, err := template.New("html").Parse(read_tmpl())
	if err != nil {
		return err
	}
	if err := index_tmpl.Execute(public_index, final); err != nil {
		return err
	}
	return nil
}

func read_tmpl() string {
	f, err := ioutil.ReadFile("./tmpl/index.html")
	if err != nil {
		log.Println("read tmpl/index.html err", err)
		return ""
	}
	return string(f)
}
