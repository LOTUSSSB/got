package greet

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

//建立一个Commit对象initCommit，message为initial commit，
//parents和blobID都为空（注意不是null，而是一个空列表，如果为null会在产生hashcode环节报错），
//date为当前时间

type Commit struct {
	Message            string
	PathToBlobID       map[string]string
	Parents            []string
	CurrentTime        time.Time
	TimeStamp          string
	ID                 string
	CommitSaveFileName string
}

func NewCommit(message string, pathToBlobID map[string]string, parents []string) *Commit {
	currentTime := time.Now()
	timeStamp := currentTime.Format(time.RFC3339)
	id := generateID(message, parents, pathToBlobID)
	commitSaveFileName := generateFileName()
	//pathToBlobID是一个map，key是文件名，value是blobID
	//parents是一个列表，存放的是当前commit的父commit的ID
	//currenttime是当前时间
	//timestamp是当前时间的格式化字符串
	//id是commit对象的ID
	//commitSaveFileName是commit对象的保存文件名

	return &Commit{
		Message:            message,
		PathToBlobID:       pathToBlobID,
		Parents:            parents,
		CurrentTime:        currentTime,
		TimeStamp:          timeStamp,
		ID:                 id,
		CommitSaveFileName: commitSaveFileName,
	}
}

func FirstInitialCommit(message string, parents []string, pathToBlobID map[string]string) *Commit {
	currentTime := time.Now()
	timeStamp := currentTime.Format(time.RFC3339)
	// 生成 commit 对象的 ID
	id := generateID(message, parents, pathToBlobID)
	commitSaveFileName := generateFileName()

	// 返回 commit 对象，创建第一条message和parents为空的commit
	return &Commit{
		Message:            "first initial commit",
		PathToBlobID:       make(map[string]string),
		Parents:            make([]string, 0),
		CurrentTime:        currentTime,
		TimeStamp:          timeStamp,
		ID:                 id,
		CommitSaveFileName: commitSaveFileName,
	}
}

func generateID(message string, parents []string, pathToBlobID map[string]string) string {
	// 生成commit对象的ID，用到message和parents，blobID
	// 用sha1加密
	h := sha1.New()
	h.Write([]byte(message))
	for _, parent := range parents {
		h.Write([]byte(parent))
	}
	for _, blobID := range pathToBlobID {
		h.Write([]byte(blobID))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// 编写generateFileName函数，内含OBJECT_DIR,ID 用于生成commit对象的保存文件名
func generateFileName() string {
	//return fmt.Sprintf("%s/%s", main.OBJECT_DIR, generateID)
	return fmt.Sprintf("%s/%s", OBJECT_DIR, generateID)
}

// 获取最新的PathToBlobID
func (c *Commit) GetPathToBlobID() map[string]string {
	return c.PathToBlobID
}

func (c *Commit) isFilePathExists(path string) bool {
	// 遍历PathToBlobID映射，如果存在与filePath相同的路径，则返回true
	for k := range c.PathToBlobID {
		if k == path {
			return true
		}
	}
	return false
}

func (c *Commit) isBlobIDExists(id string) bool {
	// 遍历PathToBlobID映射，如果存在与id相同的blobID，则返回true
	for _, v := range c.PathToBlobID {
		if v == id {
			return true
		}
	}
	return false
}
