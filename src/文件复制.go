package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 定义源文件
	sourceFilePath := "./src/example_write.txt"
	// 定义目标文件
	targetFilePath := "./src/example_write_copy.txt"
	// 对文件进行读取
	content, readErr := ioutil.ReadFile(sourceFilePath)
	if readErr != nil {
		fmt.Println("文件读取失败", readErr)
	}
	// 对文件进行写入
	writeErr := ioutil.WriteFile(targetFilePath, content, 0666)
	if writeErr != nil {
		fmt.Println("文件写入失败", writeErr)
	}
	fmt.Println("文件复制成功")
}