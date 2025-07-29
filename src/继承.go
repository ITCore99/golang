// Go语言并不是一种传统意义上的面向对象编程语言，他没有类和继承的概念
// Go使用结构体和接口来实现类似的功能
// Go 语言没有传统面向对象的类和继承的概念而是通过组合和接口来实现类似的功能。
// 通过在子结构体中加入父级的匿名结构体来实现继承的功能的特性 
package main

import (
	"fmt"
)

// 组合 Composition
// 组合是Go实现代码复用的主要方式。通过将一个结构体嵌入到另一个结构体中，子结构体可以继承父结构体的字段和方法

// // 父结构体
// type Animal struct {
// 	Name string
// }
// // 父结构体方法
// func (a *Animal) Speak() {
// 	fmt.Println(a.Name, "says hello!")
// } 

// // 子结构体
// type Dog struct {
// 	Animal // 嵌入 Animal 结构体
// 	Breed string
// }

// 接口
// 接口是Go中实现多态的主要方式，通过定义接口，不同结构体可以实现相同的方法，从而实现类似继承的多态行为。

// 定义接口
type Speaker interface {
	Speak() 
}

// 父结构体
type Animal struct {
	Name string
}
// 实现接口的方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
	Animal
	Breed string
}

func main() {
	// dog := Dog{
	// 	Animal: Animal{Name: "Buddy"},
	// 	Breed: "Golden Retriever",
	// }
	// dog.Speak() //调用父结构体方法
	// fmt.Println("Breed", dog.Breed)

	var speaker Speaker
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed: "Golden Retriever",
	}
	speaker = &dog // 把*Dog类型赋值给接口变量speaker。s = dog  ❌ 错误，因为 dog 是 Dog，不是 *Dog，不能调用 *Animal 的 Speak 方法
	speaker.Speak()
}