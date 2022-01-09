# diary

给自己写的小日记本

# 设计思路

```
git commit --allow-empty
git log --all --no-merges --shortstat --reverse
```

https://github.com/lesnitsky/git-tutor

guoyi 说她要写完之后同步到 `git` 备份，不如直接在 `git commit` 写，再从 `git log` 里面读取数据生成网页

不同分支应该有不同的处理，但当前不处理

# 主题参考

```
https://github.com/yihui/hugo-xmin
```

# 编译

主要参考 `https://github.com/libgit2/git2go#main-branch-or-vendored-static-linking`

1. `go import` 后会自动下载相应的包到 `gopath/pkg`
2. 把 `git2go` 复制出来
3. 编译 C 代码 `make install-static` （ `libgit2` 的编译也需要安装一些软件，看这里 `https://libgit2.org/docs/guides/build-and-link/` ）
4. `go mod` 修改 `replace`
5. `go build -tags static diary.go`