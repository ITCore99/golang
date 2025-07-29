package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func factional(n int, total int) int {
	if n == 0 || n == 1 {
		return total
	}
	return factional(n - 1, n * total)
}

func factional2(n int) int {
	total := 1
	if n == 0 || n == 1 {
		return total
	}

	for i := 1; i <= n; i++ {
		total *= i
	}

	return total
}

func fibonacci(n int) int {
	if n < 2  {
		return n
	}
	return fibonacci(n - 2) + fibonacci( n - 1)
}

func fibonacci2(n int) int {
	if n < 2 {
		return n
	}
	bp := []int{1, 1} // 定义一个slice
	for i := 2; i < n; i++ {
		bp = append(bp, bp[i - 2] + bp[i - 1])
	}
	// for _, value := range bp {
	// 	fmt.Println("bp的值", value)
	// }

	fmt.Printf("bp的值: %+v\n", bp)
	return bp[len(bp) - 1]
}

func walkDir(dir string, indent string) {
	entries, err := os.ReadDir(dir)
	if err != nil { // 读取路径报错
		return
	}
	fmt.Printf("目录开始读取: %s\n", dir)
	for _, entry := range entries {
		fmt.Println(indent + entry.Name())
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), indent + "  ")
		}
	}
}



func main() {
	// fmt.Println("递归函数")
	// fmt.Println(factional(5, 1))
	// fmt.Println("迭代版本", factional2(5))
	// fmt.Println("斐波那契", fibonacci(10))
	// fmt.Println("斐波那契迭代", fibonacci2(10))
	walkDir("C:/Users/86176/Desktop/front_projects/ccip_fe/dist", "")
}

