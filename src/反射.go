// 反射可以做什么
// (1) 反射可以在运行时动态获取变量的各种信息，比如变量类型，类别等信息
// (2) 如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段、方法)
// (3) 通过反射，可以修改变量的值，可以调用关联的方法
// (4) 使用反射，需要import "reflect"
// 反射相关函数
// (1) reflect.TypeOf(变量名) 获取变量的类型，返回reflect.Type类型
// (2) reflect.ValueOf(变量名) 获取变量名的值，返回reflect.Value类型(reflect.Value是一个结构体类型)通过Reflect.Value可以获取到关于该变量的很多信息
// 变量类别(kind)和类型(type)
package main 

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

// 利用一个函数返回值为空接口
func testReflect(i interface {}) {
	// 调用TypeOf 返回reflect.Type 类型结构
	reflectType := reflect.TypeOf(i)
	fmt.Println("reflectType", reflectType)
	// 调用ValueOf 返回reflect.Value 类型结构
	reflectValue := reflect.ValueOf(i)
	fmt.Println("reflectValue", reflectValue)
}

func main() {
	// 对基本类型进行反射
	var num int = 100
	testReflect(num)
	stu := Student{
		Name: "学生"
		Age: 18
	}
	fmt.Println("反射")
}