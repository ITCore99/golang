package main 
import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 连接用完一定关闭
	defer conn.Close()
	for {
		// 创建一个切片 将读取的数据全部放到切片中
		buf := make([]byte, 1024)
		// 从conn连接中读取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("连接读取数据失败", err)
			return
		}
		fmt.Println("读取客户端消息成功:", string(buf[0: n]))
	}
}

func main() {
	fmt.Println("服务器端启动")
	// 启动服务端监听
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("服务端启动失败", err)
		return
	}
	// 循环等待客户端连接
	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("等待客户端的连接失败", err2)
		} else {
			fmt.Printf("等待链接成功. con=%v,  接受到的客户端信息是: %v \n", conn, conn.RemoteAddr().String())
		}
		// 准备一协程 协程处理客户端服务的请求
		go process(conn)
	}
}