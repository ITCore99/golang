// 切片 
// Go 语言中切片是对数组的抽象。 Go的数组的长度不可改变，在特定的中这样的集合不太适用，Go中提供了一种灵活，功能强悍的内置类型切片(动态数组)与数组相比切片的长度是不固定的可以追加再追加时候
// 可能使切片的容量增大。
// go 语言中数组属于值类型。 属于值传递，如果要直接修改需要转递指针。
package main
import "fmt"

func main() {
	// var slice1 []type
	// var slice1 []type = make([]type, len)
	// slice1 := make([]type, len)
	// slice1 := make([]type, length, capacity) // 这里len是数组长度也是切片的初始长度

	// 直接初始化切片 []表示切片类型，{1, 2, 3}初始化值依次是1 2 3 其中cap=len=3
	// s :=[]int{1, 2, 3}
	// n :=  123
	// var m int = 456
	// l := int 789 (错误写法)  变量名 := 表达式
	arr := []byte{'x', 'y'} // 声明一个数组 申明一个空切片 arr := []byte{} 或者 arr := make([]byte,0) 只要[]中没有声明长度的都是切片如果要声明数组 arr := [3]byte{'1', '2', '3'}必须要指定长度
	arr2 := make([]byte, 0)
	var arr2 [3][2]int16 = [3][2]int16{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Printf("%v\n", arr)
	
 
	var numbers  = make([]int, 3, 5)
	var numbers2 []int
	printSlice(numbers)

	// 一个切片在未初始化之前默认为nil 长度为0 
	if(numbers2 == nil) {
		fmt.Printf("切片是空的")
	}

	// 截取切片
	numbers3 := []int{0, 1,2, 3, 4,5,6,7,8,10}
	fmt.Println("number3 ===", numbers3)
	fmt.Println("number3[1:4] => ", numbers3[1:4])
	numbers3  = append(numbers3, 11)
	fmt.Println("number3 append后 ===", numbers3)


}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x),x)
}