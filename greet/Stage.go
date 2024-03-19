package greet

import "os"

func readAddStage() interface{} {
	//如果addstage文件不存在，则打印addstage file not exist
	if _, err := os.Stat(ADDSTAGE_FILE); os.IsNotExist(err) {
		return ("addstage file not exist")
	}
	return nil
}
