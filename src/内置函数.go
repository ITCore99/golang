// 存在于buildin的包下面
// 1、len: 获取字符串、切片、数组长度。按照字节统计。汉字占用3个字节
// 2、rune: 字符串遍历。试讲字符串转为为了切片再对切片进行遍历. 
// 字符串遍历方式：普通的for循环按照字节、使用for range、使用rune
str := "hello golang"
r := []rune(str) 
for i = 0; i < len(r); i++ {
	fmt.Printf("%c\n", r[i])
}
// 3、strconv.Atoi: 字符串转int
// 4、strconv.Itoa: 整数转字符串
// 5、strings.Contains: 查找子串是否在指定字符串中
// 6、strings.Count: 查找子串的个数
// 7、strings.Index: 第一次出现子串的索引