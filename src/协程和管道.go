// 协程: 又称微线程 是一种用户态的轻量级线程。go中所有的并发操作都是使用协程来进行实现。占用资源更少
// 作用：在执行A函数的时候，可以随时的中断，去执行B函数，然后中断继续执行A函数(可以自动切换)，注意这一切换过程并不是函数调用(没有调用语句)
// 过程像很多线程，然而协程中只有一个线程在执行（协程的本质是单线程）协程属于程序级别的切换，操作系统完全感知不到 因此更加轻量级。
// awaitGroup 用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每一个被等待的线程在结束时应调用Done方法，同时主线程里可以调用wait方法阻塞至所有线程结束。解决主线程在子协程结束后自动结束
// 主死从随
// 编写一个程序 完成如下功能
// (1) 在主线程中，开启一个协程goroutine, 该goroutine每隔1秒输出hello golang。
// (2) 在主线程中也每隔l秒输出hello msb 输出10次 退出程序。
// (3) 要求主线程和goroutine同时执行。
package main 
import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <=10; i++ {
		fmt.Println("hello golang +", strconv.Itoa(i))
		// 阻塞1秒
		time.Sleep(1 * time.Second)
	}
}

func main () {
	go test() // 开启协程 使用关键词 "go"
	for i := 1; i <=10; i++ {
		fmt.Println("hello msb +", strconv.Itoa(i))
		// 阻塞1秒
		time.Sleep(1 * time.Second)
	}
}