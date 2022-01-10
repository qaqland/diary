package main

import (
	"io"
	"io/fs"
	"log"
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
	if _, err := os.Stat("tmpl"); err != nil {
		if os.IsNotExist(err) {
			log.Println("主题模板不存在", err)
			return
		}
		log.Println(err)
		return
	}
	if _, err := os.Stat(Public); err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
			return
		}
		if err := os.Mkdir(Public, os.ModePerm); err != nil {
			log.Println("创建发布文件夹失败", err)
			return
		}
	}

	if _, err := os.Stat("tmpl/static"); err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Println(err)
		return
	}

	if err := filepath.WalkDir("tmpl/static", walkDirFunc); err != nil {
		log.Println(err)
		return
	}
}

func walkDirFunc(path string, d fs.DirEntry, err2 error) (err3 error) {
	err3 = err2

	// 准备 static 复制到 Public 下
	new_path := strings.ReplaceAll(path, "tmpl/static", Public)
	// 这里有个 bug ，Public 应该和 static 无关。已修复
	if path == "tmpl/static" {
		return
	}
	log.Println(path, "=>", new_path)

	// 文件夹复制
	if d.IsDir() {
		if _, err := os.Stat(new_path); err != nil {

			// 不是文件夹不存在的错误
			if !os.IsNotExist(err) {
				log.Println(err)
				return
			}

			// 是文件夹不存在的错误
			if err := os.Mkdir(new_path, os.ModePerm); err != nil {
				log.Println(err)
				return
			}

			return
		}
		return
	}

	// 文件复制 从网上抄的 qaq
	file1, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	file2, err := os.OpenFile(new_path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	defer file1.Close()
	defer file2.Close()
	if _, err := io.Copy(file2, file1); err != nil {
		log.Println(err)
		return
	}
	return
}
