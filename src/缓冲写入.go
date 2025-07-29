// 读取文件的内容并显示在终端(带缓存区)的方式 适合读取比较大的文件，使用os.OpenFile，file.Close, bufio.NewReader(), reader.ReadStreamg方法。
package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, err := os.OpenFile("./src/example_write.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	// 写入文件操作： --> IO 流 ---> 缓冲输出流(带缓冲区)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		lineStr := fmt.Sprintf("%d hello Golang\n", i)
		writer.WriteString(lineStr)
	}
	//流带缓存区 只有缓存区满了之后才会写入文件 等写入完成之后需要需要刷新缓存区。Flush方法将缓冲中的数据写入下层的io.Writer接口。
	writer.Flush()
	fmt.Println("文件写入成功")
}