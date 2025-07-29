package main
import  "fmt"
import "unsafe"

	const  (
		a2 = "abc"
		b2 = len(a2)
		c2 = unsafe.Sizeof(a2)
	)
	var itVal2  int = 200

// 注意go中要使用双引号
func  main () {
	fmt.Println("Hello World!")
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s" // %d表示整型数字 %s表示字符串
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	area := b * c
	fmt.Printf("面积为: %d\n", area)

	// var a *int
	// var b1 []int
	// var c1 map[string] int
	// var d1 chan int
	// var e1 func(string) int
	// var f1 error 

	intVal  := 100  // 这种写法只能通过编译自动推断类型无法显示声明类型，需要显式声明类型的话只能通过var变量声明。
	fmt.Println(intVal)
	fmt.Println(itVal2)

	fmt.Println(a2, b2, c2)

}