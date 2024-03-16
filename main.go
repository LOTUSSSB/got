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

	repo := Repository{}

	firstArg := os.Args[1]
	switch firstArg {
	case "init":
		if !validArgs(os.Args, 2) { //检查参数是否正确
			os.Exit(0)
		}
		//调用init函数
		repo.init()

	case "add":
		if !validArgs(os.Args, 3) {
			os.Exit(0)
		}

	}

}

func validArgs(args []string, num int) bool {
	if len(args) != num {
		fmt.Println("不正确的操作数")
		os.Exit(0)
	}
	return true
}
