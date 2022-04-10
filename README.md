# diary
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fqaqland%2Fdiary.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fqaqland%2Fdiary?ref=badge_shield)


给自己写的小日记本

# 设计思路

```
git commit --allow-empty
git log --all --no-merges --shortstat --reverse
```

guoyi 说她要写完之后同步到 git 备份，不如直接在 `git commit` 写，再从 `git log` 里面读取数据生成网页

不同分支应该有不同的处理，但当前不处理

# 使用

把二进制可执行文件（程序） `diary` 和主题模板 `tmpl/` 放到目标 git 仓库，执行程序

```
Usage of diary:
  -o string
        dir for export (default "./public")
        目标文件夹，可以改称 doc 啥的
        强迫症可以每次使用前手动清理
  -p int
        pagination (default 20)
        每个分页的 commit message 数量
```

需要定制什么直接修改主题模板文件

# 编译

需要 `CGO` 主要参考

- [`libgit2`](https://libgit2.org/docs/guides/build-and-link/)
- [`git2go`](https://github.com/libgit2/git2go#main-branch-or-vendored-static-linking)

以下为 linux x86_64 

```
1. go import 后会自动下载相应的包到 {gopath}/pkg/xxx
2. 把 git2go 复制出来，在有 gomod 的那一层创建文件夹 vendor/libgit2/[] libgit2 的源码放在[]
3. go.mod 这一层执行 make install-static
4. 本项目的 go.mod 修改 replace
5. 本项目的 go.mod 同级处执行 go build -tags static *.go
```

# 主题模板

`css js img` 等静态资源要放在 `tmpl/static/`

主题要命名为 `index.html`

```golang
type Fianl struct {
	Final_commit []*git.Commit
	Here         int
	Len          int
}
```
`git.Commit` 的更多信息看 [`git2go`](https://pkg.go.dev/github.com/libgit2/git2go/v31@v31.7.4?utm_source=gopls#Commit) 的文档

# 主题参考

```
https://github.com/yihui/hugo-xmin
```

# 待处理

暂无，有问题快来发 issue

qaq

---

Copyright (c) 2022 王小明

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fqaqland%2Fdiary.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fqaqland%2Fdiary?ref=badge_large)