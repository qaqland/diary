package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// 1. 创建一个 public 发布文件夹
// 2. 把 tmpl 中的静态文件复制过去
// 3. 读取本地的 tmpl 当作模板
// 4. 生成首页的 index
// 5. 暂时不考虑 markdown 与 自定义配置，如有需要直接修改模板主题

func start_public() {
	err := filepath.WalkDir("tmpl/static", walkDirFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func walkDirFunc(path string, d fs.DirEntry, err error) error {
	new_path := strings.ReplaceAll(path, "tmpl/static", Public)
	// fmt.Println(new_path)

	// 文件夹复制
	if d.IsDir() {
		exist_mk(new_path)
		return err
	}

	// 文件复制
	file1, err2 := os.Open(path)
	if err2 != nil {
		fmt.Println(err2)
		return err
	}
	file2, err2 := os.OpenFile(new_path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
		return err
	}
	defer file1.Close()
	defer file2.Close()
	_, err2 = io.Copy(file2, file1)
	if err2 != nil {
		fmt.Println(err2)
		return err
	}
	return err
}

func exist_mk(name string) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(name, 0777)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		fmt.Println(err)
		return
	}
}
