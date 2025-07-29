// Range for循环的range格式可以对slice map 数组 和字符串等进行迭代循环。
package main
import "fmt"

var pow = []int{1, 2, 4, 6, 8, 16}
// pow := []int{1, 2, 4, 6, 8, 16} //函数外面不能这样定义 只能用于函数内部
func main() {
	for key, value := range pow {
		fmt.Printf("2**%d = %d\n", key, value)
	}

	for key, value := range "hello" {
		fmt.Printf("index: %d char: %c \n", key, value)
	}

	// 创建一个空的map key是int类型 value是float32 类型
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
  map1[3] = 3.0
  map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("key is :%d - value is :%f\n",key, value)
	}
}