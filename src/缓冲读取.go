// 读取文件的内容并显示在终端(带缓存区)的方式 适合读取比较大的文件，使用os.OpenFile，file.Close, bufio.NewReader(), reader.ReadStreamg方法。
package main 

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	file, err := os.Open("./src/example.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	// 创建文件读取器
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读取到一个换行就结束 这里'\n'表示读到已这个字符为结尾的时候就结束，双引号表示字符串故这里不能用
	   if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println("读取文件出错:", err)
					return
				}
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取成功，并全部读取完毕")

}