package main

import  (
	"fmt"
	"os"
	"log"
	// "bufio"
	// "io/ioutil"
)

func main() {
	//创建文件
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println("文件创建成功")

	//按行读取
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close() // defer 延迟调用机制 defer关键词 延迟执行指定函数 直到当前函数或者方法返回之前执行
	// // 生成一个扫描器
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Printf("Error Reading file", err)
	// }


	// 一次性写入
	// content := []byte("Hello World!")
	// wirteErr := ioutil.WriteFile("output.txt", content, 0644)
	// if wirteErr != nil {
	// 	fmt.Println("Error Writing file", wirteErr)
	// 	return
	// }
	// fmt.Println("File written successfully")

	// 文件追加
	// file, err := os.OpenFile("output.txt", os.O_APPEND | os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("Error opening file", err)
	// 	return
	// }
	// defer file.Close()
	// if _, err := file.WriteString("Appended text\n"); err != nil {
	// 	fmt.Println("Error appending to file:" , err)
	// 	return
	// }
	// fmt.Println("Text appended successfully")

	//删除文件
	// err := os.Remove("test.txt")
	// if err != nil {
	// 	fmt.Println("Error deleting file：", err)
	// 	return
	// }
	// fmt.Println("File deleted successfully")

	//获取文件信息
	// fileInfo, err := os.Stat("example.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("文件名:", fileInfo.Name())
	// fmt.Println("文件大小:", fileInfo.Size())
	// fmt.Println("权限:", fileInfo.Mode())
	// fmt.Println("最后修改时间:", fileInfo.ModTime())
	// fmt.Println("是否是目录:", fileInfo.IsDir())

	entries, err := os.ReadDir(".")
	if err != nil {
					log.Fatal(err)
	}
	
	for _, entry := range entries {
					info, _ := entry.Info()
					fmt.Printf("%-20s %8d %v %v\n",
									entry.Name(),
									info.Size(),
									info.ModTime().Format("2006-01-02 15:04:05"), entry.IsDir())
		
	}
}