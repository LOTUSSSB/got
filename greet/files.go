package greet

import (
	"bufio"
	"fmt"
	"os"
)

// 接受一个文件名，返回一个File对象，对象可与写入很多地方
func WriteContents(filePath string, contents ...interface{}) error {
	// 打开文件，使用读写模式，文件权限为0644
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	// 创建一个写入缓冲区
	str := bufio.NewWriter(file)

	for _, obj := range contents {
		switch v := obj.(type) {
		case []byte:
			if _, err := str.Write(v); err != nil {
				return err
			}
		case string:
			if _, err := str.WriteString(v); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", v)
		}
	}

	// 刷新缓冲区
	if err := str.Flush(); err != nil {
		return err
	}

	return nil
	//缓冲区作用：优化文件写入的性能，减少系统调用次数
}
