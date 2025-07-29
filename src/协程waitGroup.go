package main 

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main () {
	//  wg.Add(5)
	for i := 0 ; i < 5; i++ {
		wg.Add(1) //开启一个协程则+1
		go func(n int) {
			defer wg.Done() // 协程结束-1
			fmt.Println(n)
			 
		}(i)
	}
	wg.Wait() // 阻塞等待 wg的计数器减为0的停止
}