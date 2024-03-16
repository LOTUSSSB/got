package got

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("请输入命令")
		os.Exit(0)
	}

	firstArg := os.Args[1]
	switch firstArg {
	case "init":
		if !validArgs(os.Args, 2) {
			fmt.Println("错误的命令")
			os.Exit(0)
		}
		//调用init函数
		Repository.init()

	}

}

func validArgs(args []string, num int) bool {
	if len(args) != num {
		fmt.Println("不正确的操作数")
		os.Exit(0)
	}
	return true
}
