package main
import (
	"fmt"
	// "strconv"
)


	// 类型转换 语法 T(value) 在类型转换中 我们必须保证要转换的值和目标接口类型之间是兼容， 否则编译器会报错.
	// 定义一个Writer接口
	// type Writer interface {
	// 	Write([]byte) (int, error)
	// }

	// // 实现Writer接口的结构体
	// type StringWriter struct {
	// 	str string
	// }

	// // 实现 Write 方法
	// func (sw *StringWriter) Write(data: []byte) (int, error) {
	// 	sw.str += string(data)
	// 	return len(data), nil
	// }

// 类型转化用于将一种数据类型的变量转换为另一种类型的变量 type_name(expression)
func main() {
	// 数值类型转化
	// fmt.Println("类型转换")
	// var a int = 10
	// var b float64 = float64(a)
	// fmt.Printf("b 的值为: %f\n", b)
	// var sum int = 17
	// var count int = 5
	// var mean float32
	// mean = float32(sum) / float32(count)
	// fmt.Printf("mean 的值为: %f\n", mean)
	// // 字符串类型之间转化
	// var str string = "10"
	// var num int
	// num, _ = strconv.Atoi(str)
	// fmt.Printf("num 的值为: %d\n", num)
	// numStr := strconv.Itoa(num)
	// fmt.Printf("numStr 的值为: %s\n", numStr)
	
	// floatStr := "3.14"
	// floatNum, err := strconv.ParseFloat(floatStr, 64)
	// if err != nil {
	// 	fmt.Printf("转换错误：", err)
	// } else {
	// 	fmt.Printf("字符串 '%s' 转为浮点型为: %f\n", floatStr, floatNum)
	// }

	// num :=  3.14
	// str := strconv.FormatFloat(num, 'f', 2, 64)
	// fmt.Printf("浮点数 %f 转化为字符串 %s\n", num, str)

	// 接口类型转化 接口类型转化有两种: 类型断言和类型转换 
	// 类型断言 语法：value.(type) 或者 value.(T)
	// var i interface {} = "Hello World"  // interface{} 表示空接口 可以表示任何类型的值。 所有的类型实现了空接口
	// str, ok := i.(string)
	// if !ok {
	// 	fmt.Printf("conversion failed")
	// } else {
	// 	fmt.Printf("'%s' is a string \n", str)
	// }

	// 创建一个StringWriter 实例并赋值给Writer接口
	// var w Writer = &StringWriter{}
	// //将Writer接口类型转换为StringWriter类型
	// sw := w.(StringWriter)
	// //修改StringWriter的字段
	// sw.str = "Hello World"
	// fmt.Println(sw.str)
	printValue(42)
  printValue("hello")
  printValue(3.14)
}

func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown type")
	}
}