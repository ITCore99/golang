// 错误处理
// go语言中追求代码优雅 没有使用try catch 而是引入 defer + recover机制处理错误
// recover: 允许程序管理代码执行过程中的panic，在defer函数中 执行recover调用回传panic调用的错误值，恢复正常执行，停止恐慌。如在defer函数之外被调用，他将不会停止恐慌过程序列
// 在此情况下，或当该程序不在恐慌过程中，或提供panic实参为nil是recover返回
package main 

import (
	"fmt"
	"errors"
	"math"
)

type error interface {
	Error() string
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

//fmt包与错误格式化

// errors.Is 和 errors.As 检查某个错误是否为指定错误是由该错误包装而成

// panic Go中panic用于处理不可恢复的的错误，recover用于从panic中恢复。

func main() {
	// err := errors.New("this is an error")
	// fmt.Println(err)
	// result, err := Sqrt(4)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Sqart 结果为：%f\n", result)
	// }

	// 测试案例1：正常情况（无错误）
	num1, num2 := 10, 2
	result, err := test(num1, num2)
	
	// 在调用位置进行判断
	if err != nil {
		fmt.Printf("计算 %d / %d 失败: %v\n", num1, num2, err)
	} else {
		fmt.Printf("计算 %d / %d 成功: 结果为 %d\n", num1, num2, result)
	}

	// 测试案例2：错误情况（除数为零）
	num3, num4 := 8, 0
	result2, err2 := test(num3, num4)
	
	if err2 != nil {
		fmt.Printf("计算 %d / %d 失败: %v\n", num3, num4, err2)
	} else {
		fmt.Printf("计算 %d / %d 成功: 结果为 %d\n", num3, num4, result2)
	}
}

func test(num1 int, num2 int) (int, error) {
	var result int
  var err error

	// 利用defer后加上匿名函数的调用 + recover 来捕获错误
	// defer func() {
	// 	// 调用recover内置函数， 将错误进行捕获panic错误
	// 	e := recover()  // 返回一个错误使用err接收。 如果没有发生错误那么err就为零值：nil
	// 	if e != nil {
	// 		fmt.Println("错误已经捕获")
	// 		result = 0
	// 		err = fmt.Errorf("%v", e) 
	// 	}
	// }()

	if num2 == 0 {
		// 方式1 主动触发恐慌 recover捕获到
		// panic("除数不能为空") 
		// 方式2 errors创建手动创建错误对象并返回，不会造成panic所以recover也不会捕获到
		return result, errors.New("除数不能为空")
	}
	
	result= num1 / num2
	return  result, err
	// return 返回计算结果和错误(此时err为nil) 在Go语言中，当使用命名返回值时，可以不用显式指定返回值，直接使用 return 即可，编译器会自动返回命名的返回值变量。 也就是返回
	// 类型限制为 (result int, err error)
}