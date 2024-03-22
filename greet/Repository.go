package greet

import (
	"os"
	"path/filepath"
)

// os.Getenv函数获取环境变量的值
// join函数将多个字符串连接成一个路径
var (
	CWD              = os.Getenv("PWD")
	GotDir           = filepath.Join(CWD, ".got") //返回相对路径
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
func (r *Repository) Init() {

	//检查是否已经初始化，若未初始化则创建文件夹
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

// 检查是否已经初始化，若未初始化则继续执行Init程序
func (r *Repository) checkIfInitialized() {
	//如果.Got文件夹已经存在，则打印已经初始化
	if _, err := os.Stat(GotDir); err == nil {
		println("Already initialized")
		os.Exit(0)
	}
	//如果.Got文件夹不存在，则继续执行init函数
	//Init()
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
	//writeContents(HEAD_FILE, "master")
	//调用files.go中的writeContents函数
	WriteContents(HEAD_FILE, "master")

}

func (r *Repository) initHeads() {
	//读取HEADS_DIR下的所有文件，返回一个文件名列表
	//如果HEADS_DIR不存在，返回一个空列表
	files, _ := filepath.Glob(HEADS_DIR + "/*")
	//如果文件列表为空，调用writeContents函数，将currentCommit.ID写入HEADS_DIR下的master文件
	if len(files) == 0 {
		WriteContents(filepath.Join(HEADS_DIR, "master"), currentCommit.ID)
	}

}

func (r *Repository) Commit(message string) {
	//调用commit.go中的NewCommit函数
	//NewCommit(message, currentCommit.PathToBlobID, []string{currentCommit.ID})
	//调用commit.go中的NewCommit函数
	NewCommit(message, currentCommit.PathToBlobID, []string{currentCommit.ID})
}

func (r *Repository) Add(filePath string) {
	//调用add.go中的add函数
	//add(filePath)

}

// 保存Blob对象（将文件的路径和文件的blobID关联并存储起来）
func storeBlob(blob *Blob) {
	// 读取当前提交、添加阶段和删除阶段
	currCommit := readCurrCommit().(*Commit)
	addStage := readAddStage()
	removeStage := readRemoveStage()

	//	// 如果当前提交中不包含该 blob 或者删除阶段中已存在该 blob，则进行存储
	//	if !currCommit.PathToBlobID.Contains(blob.BlobID) || !removeStage.isNewBlob(blob) {
	//		if addStage.isNewBlob(blob) {
	//			if removeStage.isNewBlob(blob) {
	//				// 如果添加阶段和删除阶段都不存在该 blob，则进行存储并更新添加阶段
	//				blob.save()
	//				if addStage.isFilePathExists(blob.Path) {
	//					addStage.delete(blob)
	//				}
	//				addStage.add(blob)
	//				addStage.saveAddStage()
	//			} else {
	//				// 如果只有删除阶段存在该 blob，则在删除阶段中删除该 blob
	//				removeStage.delete(blob)
	//				removeStage.saveRemoveStage()
	//			}
	//		}
	//	}

	//// 如果当前提交中不包含该 blob 或者删除阶段中已存在该 blob，则进行存储
	//if !currCommit.PathToBlobID.Contains(blob.BlobID) || !removeStage.isNewBlob(blob) {
	//	if addStage.isNewBlob(blob) {
	//		if removeStage.isNewBlob(blob) {
	//			// 如果添加阶段和删除阶段都不存在该 blob，则进行存储并更新添加阶段
	//			blob.save()
	//			if addStage.isFilePathExists(blob.Path) {
	//				addStage.delete(blob)
	//			}
	//			addStage.add(blob)
	//			addStage.saveAddStage()
	//		} else {
	//			// 如果只有删除阶段存在该 blob，则在删除阶段中删除该 blob
	//			removeStage.delete(blob)
	//			removeStage.saveRemoveStage()
	//		}
	//	}
	//}

	//大概写个思路？ 感觉go不能一次性这样来，要不停的调用函数

}

func readCurrCommit() interface{} {
	//调用repository.go中的	currentCommit *Commit 指针
	return currentCommit
}

// 判断Blob是否为最新的
func isNewBlob(blob *Blob) bool {
	//调用blob.go中的isNewBlob函数
	return isNewBlob(blob)
}
