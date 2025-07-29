// select功能： 解决多个管道的选择问题，也可以叫做多路复用，可以从多个管道中随机公平地选择一个来来执行。
// ps: case后面必须进行的是io操作，不能是等值，随机去选择io操作
// ps: default防止select被阻塞住，加入default
package main 

import (
	"fmt"
	"time"
)

func main() {
	// 定义int类型管道
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 5)
		intChan<- 10
	}()
	// 定义string类型管道
	stringChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 12)
		stringChan<- "hello Golang"
	}()
	fmt.Println(<-intChan) // 本身取数据就是阻塞的

	select {
	case v:= <-intChan: // case 后面必须是IO操作
		fmt.Println("intChan:", v)
	case v:= <-stringChan: 
		fmt.Println("stringChan:", v)
	default:
		fmt.Println("防止select被阻塞")
	}
}