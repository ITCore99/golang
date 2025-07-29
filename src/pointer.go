package main
import "fmt"

const MAX int = 3
func main() {
	a := [3]int{10,100,200}
	var i int
	//有一种情况我们可能需要保存数组，这样我们就需要使用到指针 以下声明了整型指针数组
	var ptr [MAX]*int;

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]  //整数地址赋值给指针数组
	}
	
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d | 地址为: %d\n", i, *ptr[i], ptr[i])
	}

	// 指向指针的指针。 访问指向指针的指针变量的值需要使用两个*号
	// var ptr **int

	var b int
	var ptr1 *int
	var pptr1 **int
	b = 3000
	
	ptr1 = &b // 指针 ptr 地址
	pptr1 = &ptr1

	fmt.Printf("变量 b = %d\n", b)
	fmt.Printf("指针变量 *ptr1 = %d\n", *ptr1)
	fmt.Printf("指针变量 **pptr1 = %d\n", **pptr1)

	// go 语言中允许向函数传递指针。只需要在函数定义的参数上设置为指针类型即可
	var c int = 100
	var d int = 200

	fmt.Printf("交换前 c 的值 : %d\n", c)
	fmt.Printf("交换钱 d 的值 : %d\n", d)

	swap(&c, &d)

	fmt.Printf("交换后 c 的值 : %d\n", c)
	fmt.Printf("交换后 d 的值 : %d\n", d)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y  // 将*y的值赋值给x所指向的内存地址中 *x = something 的意思是：将 something 的值 写入到 x 指向的内存地址中。
	*y = temp
}