// case 后面不需要写 break
// case 后面可以写多个值，多个值之间以逗号分隔。
// default 位置不是固定的可以写在前面
// switch穿透 利用fallthrough关键字，如果在case语句块后增加fallthrough，则会继续执行下一个case。也叫switch穿透。
// switch 后面可以不带表达式，当做if分支来使用
// switch {
// 	case score == 100 :
// 			fmt.Println("分数是", 100)
// 	case score == 98:
// 			fmt.Println("分数是", 98)
// }

package main

import (
	"fmt"
)

func main() {
	var score int = 98
	fmt.Println("计算结果", score / 10)
	switch score / 10 {
	case 10, 9:
		fmt.Println("您的成绩为A")
	case 8:
		fmt.Println("您的成绩为A-")
	default:
		fmt.Println("您的成绩未匹配到")
	}

	switch {
		case score == 100 :
				fmt.Println("分数是", 100)
		case score == 98:
				fmt.Println("分数是", 98)
	}

	switch n := 6; { //注意：这种语法必须要以分号结尾
		case n > 3:
			fmt.Println("大雨3")
		case n < 3:
			fmt.Println("小于3")
	}
}