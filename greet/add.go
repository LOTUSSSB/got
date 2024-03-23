package greet

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type Blob struct {
	ID               string
	Bytes            []byte
	FileName         string
	FilePath         string
	BlobSaveFileName string
}

func (b Blob) Save() {
	WriteContents(b.BlobSaveFileName, b.Bytes)
}

//func generateID(message string, parents []string, pathToBlobID map[string]string) string {
//	// 生成commit对象的ID，用到message和parents，blobID
//	// 用sha1加密
//	h := sha1.New()
//	h.Write([]byte(message))
//	for _, parent := range parents {
//		h.Write([]byte(parent))
//	}
//	for _, blobID := range pathToBlobID {
//		h.Write([]byte(blobID))
//	}
//	return hex.EncodeToString(h.Sum(nil))
//}
// 创建一个Blob对象

func NewBlob(file string) *Blob {
	fileName, _ := filepath.Abs(file)
	bytes := readFile(fileName)
	filePath := filepath.Dir(fileName)
	//需要传入绝对路径来工作
	//id := generateID(filePath, bytes)
	//之前的generateID函数是用来生成commit对象的ID的，这里需要一个新的函数来生成blob对象的ID😅
	id := generateBlobId(fileName)
	//Q: 这里生成的ID是什么？,与commitID有什么区别？
	//A: 这里的ID是blob对象的ID，是根据文件内容生成的唯一标识符
	blobSaveFileName := generateBlobSaveFileName(id)

	return &Blob{
		ID:               id,
		Bytes:            bytes,
		FileName:         filepath.Base(fileName),
		FilePath:         filePath,
		BlobSaveFileName: blobSaveFileName,
	}
}

func generateBlobSaveFileName(id string) string {
	return fmt.Sprintf("%s/%s", OBJECT_DIR, id)
}

//// 保存Blob对象
//func add(file string) {
//	fileName, _ := getFileFromCWD(file)
//	//if !exists(fileName) {
//	//	fmt.Println("File does not exist.")
//	//	os.Exit(0)
//	//}
//	//传入文件名查看文件是否存在
//	//如果文件不存在，打印文件不存在并退出
//	if _, err := os.Stat(fileName); os.IsNotExist(err) {
//		fmt.Printf("文件 %s 不存在\n", fileName)
//		os.Exit(1)
//	}
//	blob := NewBlob(fileName)
//	storeBlob(blob)
//}

// Add 将blob和文件路径FilePath关联并存储起来
func (s *Stage) Add(blob Blob) {
	s.PathToBlobID[blob.FilePath] = blob.ID
}

// Delete 删除blob和文件路径的关联

// 获取文件的绝对路径
func getFileFromCWD(file string) (string, error) {
	absPath, _ := filepath.Abs(file)
	return absPath, nil
}

//// 检查文件是否存在
//func exists(file string) bool {
//	// 检查文件是否存在
//}

// 传入文件名，生成blob对象的ID
func generateBlobId(fileName string) string {
	// 生成blob对象的ID，用到文件名和文件内容
	// 用sha1加密
	h := sha1.New()
	h.Write([]byte(fileName))
	bytes := readFile(fileName)
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

// readFile函数，传入文件名，返回文件内容
func readFile(fileName string) []byte {
	// 读取文件内容
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	fi, _ := file.Stat()
	size := fi.Size()
	bytes := make([]byte, size)
	_, err = file.Read(bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytes
}
