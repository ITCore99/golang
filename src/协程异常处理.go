// 多个协程工作，其中一个协程出现panic, 导致程序崩溃
// 解决方法: 利用defer + recover捕获panic进行处理，即使协程出问题，主线程仍然不受影响可以继续工作，
package main 

import (
	"fmt"
	_"time"
	"sync"
)

var wg sync.WaitGroup
// 输出数字
func printNum() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
// 做除法操作
func devide() {
	defer func () {
		wg.Done()
		err := recover()  //使用recover捕获panic
		if err != nil {
			fmt.Println("devide出现错误", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println("devide 除法",result)
}

func main() {
	wg.Add(2)
	//启动两个协程
	go printNum()
	go devide()

	wg.Wait()

}