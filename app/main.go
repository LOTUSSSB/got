package main

import (
	"fmt"
	"main/greet"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("请输入命令")
		os.Exit(0)
	}

	repo := greet.Repository{}

	firstArg := os.Args[1]
	switch firstArg {
	case "init":
		if !validArgs(os.Args, 1) { //检查参数是否正确
			os.Exit(0)
		}
		//调用init函数
		repo.Init()

	case "add":
		if !validArgs(os.Args, 2) {
			os.Exit(0)
		}
		//调用add函数
		repo.Add(os.Args[2])

	case "commit":
		if !validArgs(os.Args, 2) {
			os.Exit(0)
		}
		repo.Commit(os.Args[2])

	case "rm":
		if !validArgs(os.Args, 2) {
			os.Exit(0)
		}
		repo.Remove(os.Args[2])

	}
}

// 生成字符串数组来检测输入的命令行参数数量是否正确
func validArgs(args []string, num int) bool {
	if len(args) != num {
		fmt.Println("不正确的操作数")
		os.Exit(0)
	}
	return true
}
