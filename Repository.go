package got

import (
	"fmt"
	"os"
	"path/filepath"
)

// os.Getenv函数获取环境变量的值
// join函数将多个字符串连接成一个路径
var (
	CWD              = os.Getenv("PWD")
	GotDir           = filepath.Join(CWD, ".got")
	OBJECT_DIR       = filepath.Join(GotDir, "objects")
	RefsDir          = filepath.Join(GotDir, "refs")
	HEADS_DIR        = filepath.Join(RefsDir, "heads")
	HEAD_FILE        = filepath.Join(GotDir, "HEAD")
	ADDSTAGE_FILE    = filepath.Join(GotDir, "add_stage")
	REMOVESTAGE_FILE = filepath.Join(GotDir, "remove_stage")
	//创建·一个currentCommit指针，指向当前的commit
	currentCommit *Commit
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

	//调用initcommit函数
	r.initcommit()
	r.initHEAD()
	r.initHeads()

}

func (r *Repository) checkIfInitialized() {
	if _, err := os.Stat(GotDir); os.IsNotExist(err) {
		fmt.Println("Not in an initialized Gitlet directory.")
		os.Exit(0)
	}
}

func (r *Repository) initcommit() {
	//initcommit := &Commit{}
	//currentCommit = initcommit
	//调用commit.go中的FirstInitialCommit函数
	currentCommit = FirstInitialCommit("first initial commit", []string{}, make(map[string]string))
}

// 将HEAD文件的内容设置为"master"。HEAD文件是版本控制系统中的一个特殊文件，
// 它指示当前所在的分支。在这个方法中，通过调用writeContents函数将"master"写入HEAD文件。
func (r *Repository) initHEAD() {
	//调用writeContents函数，将"master"写入gotdir下的HEAD文件
	writeContents(HEAD_FILE, "master")

}

func (r *Repository) initHeads() {
	//读取HEADS_DIR下的所有文件，返回一个文件名列表
	//如果HEADS_DIR不存在，返回一个空列表
	files, _ := filepath.Glob(HEADS_DIR + "/*")
	//如果文件列表为空，调用writeContents函数，将currentCommit.ID写入HEADS_DIR下的master文件
	if len(files) == 0 {
		writeContents(filepath.Join(HEADS_DIR, "master"), currentCommit.ID)
	}

}
