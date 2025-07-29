// 通过对多个协程加上互斥锁实现
package main
import (
	"fmt"
)

func main() {
	fmt.Println("多协程操作同一个数据")
}