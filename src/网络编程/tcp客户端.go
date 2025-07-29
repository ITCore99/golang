// 调用Dial
package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	_"time"
)

func main() {
	fmt.Println("客户端启动...")
	// 调用Dial
	conn, err := net.Dial("tcp", "127.0.0.1:3001")
	if err != nil {
		fmt.Println("客户端启动发生错误", err)
		return
	}
	fmt.Println("链接成功", conn)

	for {
		reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准的终端(比如键盘等)输入
		str, err :=reader.ReadString('\n')
		if err != nil {
			fmt.Println("终端输入失败", err)
		} else {
			n, err := conn.Write([]byte(str))
			if err != nil {
				fmt.Println("发送数据失败", err)
				return
			}
			fmt.Printf("终端数据通过客户端发送成功，一共发送了%d个字节\n", n)
		}
	}
}