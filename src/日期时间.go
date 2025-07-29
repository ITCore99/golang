package main
 
import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())

	// Sprintf 可以得到固定格式化的字符串 以便后续使用
 dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
 fmt.Println(dateStr)
 fmt.Println("格式化时间", now.Format("2006-01-02 15:04:05")) // 将时间格式化为指定格式的字符串
}