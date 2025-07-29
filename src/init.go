package main 
// init函数 初始化函数 它会先于main函数进行执行  进行一些初始化工作
// 每一个源文件都可以包含一个init函数。对于多个init文件时被导入的init函数会先执行。
// 全局变量 init函数 main函数的执行顺序是:全局变量 > init > main
import (
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	// new 分派内存(主要是值类型) 内置函数返回的是一个类型的的指针
	num := new(int) 
	fmt.Printf("num的类型为: %T, num的值为: %v, num的地址为: %v, num指针指向的值为: %v", num, num, &num, *num)
}