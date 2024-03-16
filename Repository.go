package got

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	CWD              = os.Getenv("PWD")
	GotDir           = filepath.Join(CWD, ".got")
	OBJECT_DIR       = filepath.Join(GotDir, "objects")
	RefsDir          = filepath.Join(GotDir, "refs")
	HEADS_DIR        = filepath.Join(RefsDir, "heads")
	HEAD_FILE        = filepath.Join(GotDir, "HEAD")
	ADDSTAGE_FILE    = filepath.Join(GotDir, "add_stage")
	REMOVESTAGE_FILE = filepath.Join(GotDir, "remove_stage")
	currentCommit    *Commit
)

/*
 *   .got
 *      |--objects
 *      |     |--commit and blob
 *      |--refs
 *      |    |--heads
 *      |         |--master
 *      |--HEAD
 *      |--addstage
 *      |--removestage
 */

//func init() {
//
//}

//创建一个Repository结构体，使得init等能被调用

type Repository struct {
}

// 创建一个init函数，使得他能被main调用，这个init函数功能是初始化got目录
func (r *Repository) init() {

	//检查是否已经初始化，若程序继续执行，表示在初始化的got目录下
	r.checkIfInitialized()

	os.Mkdir(GotDir, 0755)
	os.Mkdir(OBJECT_DIR, 0755)
	os.Mkdir(RefsDir, 0755)
	os.Mkdir(HEADS_DIR, 0755)
	os.Create(HEAD_FILE)
	os.Create(ADDSTAGE_FILE)
	os.Create(REMOVESTAGE_FILE)

}

func (r *Repository) checkIfInitialized() {
	if _, err := os.Stat(GotDir); os.IsNotExist(err) {
		fmt.Println("Not in an initialized Gitlet directory.")
		os.Exit(0)
	}
}

func (r *Repository) initcommit() {
	initcommit := &Commit{}
	currentCommit = initcommit
}

func (r *Repository) inHEAD() {

	//读取HEAD文件，返回HEAD指向的分支

}
