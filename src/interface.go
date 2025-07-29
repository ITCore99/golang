// 接口是go语言中的一种类型，用于定义行为集合，他通过描述类型必须实现的方法，规定了类型的行为契约。
// 只要一个类型实现了接口的所有方法，该类型就自动实现了该接口 
// 一个未初始化的接口其值为nil零值 
// 定义为interface {} 可以表示任何类型
// 接口本身是不能直接创建实例的只能通过被结构进行继承。
// 多态：定义一个函数：专门用来各国人打招呼的函数 接受具备(实现)SayHello接口能力的变量或者结构体
// func greet(s SayHello) {
// 	s.sayHello()
// }

package main

import (
	"fmt" 
	"math"
)


type  Shape interface {
	Area() float64
	Perimeter() float64
}

//定义一个结构体
type Circle struct {
	Radius float64
}

//Circle实现Shape接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 定义一个矩形结构体
type Rectangle struct {
	Width float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// 空接口 interface{} 是Go的特殊接口 表示所有类型的超集
func printValue(val interface{}){
	fmt.Printf("Value: %v, Type: %T\n",val, val)
}

// 类型断言用于从接口类型中提取其底层值 value := value.(Type)

// 带类型检查的断言 为了避免panic 可以使用带检查的类型断言

// 类型选择 
// type switch是Go中的语法结构用于更具变量的类型执行不同的逻辑结构
func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Interge: ", v)
	case string:
		fmt.Println("String", v)
	case float64:
		fmt.Println("Float", v)
	default:
		fmt.Println("UnKnown type")
	}
}

type  Reader interface {
	Read() string
}
type Writer interface {
	Write(data string) 
}

type ReadWriter interface {
	Reader
	Writer
}

// 定义一个结构体
type File struct {
}

func (f File) Read() string {
	return "Reading data"
}
func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
} 


type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func main() {
	c := Circle{Radius: 5}
	var s Shape = c   // 接口变量可以存储实现了接口的类型
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
	printValue(42)
	printValue("hello")
	printValue(3.14)
	printValue([]int{1, 2})

	var i interface{} = "Hello World"
	str := i.(string) // 类型断言
	fmt.Printf("Value: '%s' Type: %T\n", str, str) 

	var i1 interface{} = 42
	str, ok := i1.(string)
	if ok {
		fmt.Println("String: ", str)
	} else {
		fmt.Println("Not a string")
	}

	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})

	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello Go!")

	var phone Phone
  phone = new(NokiaPhone)  // new 是Go中一个内置函数 new(T) 会创建一个T类型的变量，返回该变量的地址 所以new(NokiaPhone)类型是*NokiaPhone
  phone.call()
  phone = new(IPhone)
 phone.call()

// 多态数组(可以存放实现了当前接口的任何数据类型)
var shapes [3]Shape
shapes[0] = Circle{Radius: 20}
shapes[1] = Rectangle{Width: 2, Height:3}
shapes[2] = Rectangle{Width: 20, Height: 30}
fmt.Println("shapes", shapes, shapes[1].Area())

}