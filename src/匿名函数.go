// Go支持匿名函数 如果我们某个函数只是希望使用一次，可以考虑使用匿名函数
// 匿名函数使用方式
// 1、在定义匿名函数时就直接调用，这种方式的匿名函数只能调用一次
// 2、将匿名函数赋值给一个变量，再通过变量名来调用匿名函数
// 3、如何让一个匿名函数可以在整个程序中有效呢？将匿名函数赋值 给一个全局变量就可以了

package main
import "fmt" 

func main() {
	res := func (num1 int, num2 int) int {
		return num1 + num2
	}(1, 2)
	fmt.Println("res", res)
}