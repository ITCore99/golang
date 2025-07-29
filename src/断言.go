package main

import (
	"fmt"
)

type SayHello interface {
	sayHello()
}

type Chinese struct {
	name string
}

func (c Chinese) sayHello() {
	fmt.Printf("%s: 你好\n", c.name)
}
func (c Chinese) xiaqi() {
	fmt.Println("下棋")
}

type American struct {
	name string
}
func (a American) sayHello() {
	fmt.Printf("%s: hello\n", a.name)
}

func greet(s SayHello) {
	// switch v := s.(type) // 这里的type是一个go中关键字
	// 进行断言转化为中国人
	// if ch, ok := s.(Chinese); !ok {
	// 	fmt.Println(ch, "不是中国人")
	// } else {
	// 	ch.xiaqi()
	// }

	switch s.(type) {
		case Chinese:
			ch := s.(Chinese)
			ch.xiaqi()
		case American:
			ch := s.(American)
			ch.sayHello()
	}
}


func main () {

	c := Chinese{name: "小明"}
	a := American{name: "Jonny"}
	c.sayHello()
	a.sayHello()
	greet(c)
	greet(a)

}