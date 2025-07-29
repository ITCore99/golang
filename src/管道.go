// 管道
// 1、管道的本质是一个队列
// 2、数据是先进先出
// 3、自身线程安全 多协程访问时不需要加锁channel本身就是线程安全
// 4、管道是有类型的， 一个string的管道只能存放string类型的数据。
// 5、管道只支持for-range的方式遍历。并且遍历时需要注意两个细节 （1）、在遍历时如果管道没有关闭，则回出现deadlock错误。（2）、在遍历时如果管道已经关闭了，则会正常遍历，遍历完后就退出遍历。
// 6、默认情况下管道是双向的。
// 当程序只往管道里面写而不往出读取的话造成管道的阻塞。只要有读的话就不会造成阻塞 即使是写入和读取的频率不一致。

package main

import (
	"fmt"
)

func main() {
	// 定义管道
	var intChan chan int
	// 通过make初始化：管道可以存放3个int类型的数据
	intChan = make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i
	}
	close(intChan)
	// 遍历
	for v := range intChan {
		fmt.Println("value = ", v)
	}

	// 声明为只写
	var intChanW chan<- int
	// 声明为只读
	var intChanR chan-> int
	
}