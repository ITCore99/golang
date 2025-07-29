// map 在使用之前一定要make，只是声明变量的话是不会分配内存的。只有当make之后才会分配空间。
// map中的键是无序的，重名的键会被覆盖。
package main 
import  "fmt"

func main() {
	// m1 := make(map[string]int)
	m2 := map[string]int{
		"apple": 1,
		"banana": 2,
		"orange": 3,
	} 

	// 获取值
	v1 := m2["apple"]
	fmt.Println("apple的值 ", v1)

	// 获取键值对
	// v1 := m2["apple"]
	// v2, ok := m2["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值\
	
	// 获取键值对
	// v1 := m2["apple"]
	// v2, ok := m2["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值

	// 删除键值对
	delete(m2, "banana")

	// 获取长度
	m2Len := len(m2)
	fmt.Println("m1的长度 ", m2Len)

	// 双层嵌套map
	a := make(map[string]map[int]string)
	a["班级1"] = make(map[int]string, 3 )
	a["班级1"][20096677] = "张三"
	a["班级1"][20096688] = "李四"
	a["班级1"][20096699] = "王五"

	fmt.Println("a", a, len(a))
}