// 案例需求
// 请完成协程和管道协同工作的案例，具体要求如下
// (1) 开启一个writeData协程，向管道中写入10个整数
// (2) 开启readData协程从管道中读取WriteData 写入的数据
// (3) 注意：writeData 和 readData操作的是同一个管道
// (4) 主线程需要等待writeData 和 readData协程都完成后才能退出


package main

import (
	"fmt"
	"sync"
	"time"
)

// 写数据
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		intChan <- i
		fmt.Println("写入数据为:", i)
		time.Sleep(time.Second)
	}
	//关闭管道
	close(intChan)
}

// 读数据
func readData(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("读出的数据为:", v)
		time.Sleep(time.Second)
	}
}

var wg sync.WaitGroup
func main() {
	intChan := make(chan int, 10)
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
}