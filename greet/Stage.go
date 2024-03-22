package greet

import (
	"encoding/gob"
	"os"
	"reflect"
)

type Stage struct {
	PathToBlobID map[string]string
}

func readAddStage() interface{} {
	//如果addstage文件不存在，则返回一个新的Stage对象
	if _, err := os.Stat(ADDSTAGE_FILE); os.IsNotExist(err) {
		return &Stage{}
	}
	//否则调用readObject函数
	result, err := readObject(os.Open(ADDSTAGE_FILE))
	if err != nil {
		return err
	}
	return result
}

func readRemoveStage() interface{} {
	//如果removestage文件不存在，则返回一个新的Stage对象
	if _, err := os.Stat(REMOVESTAGE_FILE); os.IsNotExist(err) {
		return &Stage{}
	}
	//否则调用readObject函数
	result, err := readObject(os.Open(REMOVESTAGE_FILE))
	if err != nil {
		return err
	}
	return result
}

// 从文件中读取对象，并且进行解码（反序列化）
func readObject(file *os.File, expectedClass interface{}) (interface{}, error) {
	// 创建一个 Decoder 以从文件中解码数据
	decoder := gob.NewDecoder(file)

	// 通过反射创建一个新的实例，用于存储解码后的数据
	result := reflect.New(reflect.TypeOf(expectedClass).Elem()).Interface()

	// 将文件中的数据解码到 result 中
	err := decoder.Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Stage) IsNewBlob(blob Blob) bool {
	for _, v := range s.PathToBlobID {
		if v == blob.ID {
			return false
		}
	}
	return true
}
