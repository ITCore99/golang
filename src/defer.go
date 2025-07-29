// defer关键字 在函数中经常需要创建、或者读取资源，为了在函数执行完毕后及时的释放资源。go的设计者提供了defer关键字。
// 在go程序中遇到defer关键字时不会立即执行defer后面的语句而是把defer后面的语句压入一个栈(单独的)中，继续执行函数后面的语句。等后面语句执行完成之后再执行defer栈中的数据。
// 遇到defer关键词会将后面的后面的语句压入栈中，同时也会将相关的值拷贝到栈中，不会随着后面程序执行而发生变化。
package main
import "fmt"

func main() {
	fmt.Println(add(30, 60))
}

func add(num1 int, num2 int) int {
	defer fmt.Println("num1", num1)
	defer fmt.Println("num2", num2)
  num1 += 90
	num2 += 50
	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}