// 结构体
// 方法是作用在指定的结构类型上，和指定的数据类型绑定 因此自定义类型都可以有方法而不仅仅是struct。
// 方法的访问范围控制规则，和函数是一样的。方法的首字母小写只能在本包访问，方法首字母大写可以在本包和其他包访问。
// 如果一个类型实现了String()方法，那么fmt.Println()默认调用这个变量的String()进行输出。
// 跨包使用结构体创建实例时首先需要结构体首字母大写，其次要通过import将当前结构体所在的的包导入到使用的包中。
package main
import (
	"fmt"
	"test/src/model" // 路径是go.mod里面的moudle + 当前package路径。如果没有go.mod在项目目录下执行go mod init 项目文件夹名
)

type Book struct {
	title string
	author string
	subject string
	book_id string
}

type Teacher struct {
	Name string
	Age int
	School string
}

type Student struct {
	Age int
}

type Person struct {
	Age int 
}

func (p Person) getAge() int { // 这里p传递的副本(值传递)而不是引用传递。如果希望在方法中改变结构体变量的值，可以通过结构体指针的方式来处理
	return p.Age
}

func (p *Person) setAge() int { // 这里p传递的副本(值传递)而不是引用传递。如果希望在方法中改变结构体变量的值，可以通过结构体指针的方式来处理
	(*p).Age = 11
	return p.Age
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Book{"Go 语言", "www.baidu.com", "Go 语言教程", "123"})
	//  也可以使用key=>value的形式
	fmt.Println(Book{title: "Go 语言", author: "www.baidu.com", subject: "Go 语言教程", book_id: "123"})
	// 忽略的字段为0或者空
	fmt.Println(Book{title: "Go 语言", author: "www.baidu.com"})

	var Book1 Book

	Book1.title = "Go 语言"
	Book1.author = "www.baidu.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = "123"

	// fmt.Println(Book1)
	// fmt.Printf("Book1 title :%s\n", Book1.title)

	printBook(Book1)

	var bookPtr *Book
	bookPtr = &Book1

	fmt.Printf("bookPtr 值为: %+v\n", *bookPtr)

	//通过new关键字创建结构体指针对象  
	tt := Teacher{"小五", 25, "清华大学"}
	printTeacher(&tt)
	fmt.Println("老师结构体外部", tt)

	pt := &Teacher{"小五", 25, "北京大学"}
	fmt.Println("&t", pt)
	 

	// 结构体是用户单独定义的类型，和其他类型进行转换时需要完全相同的字段(名字、个数、类型)。
	// 结构体进行type重新定义(相当于取别名)， Golang认为是新的类型，但是可以相互间强转。
	var s Student = Student{10}
	var p Person = Person{10}
 	s = Student(p) // 进行强制类型转换
	fmt.Println(s)
	fmt.Println("Person 年龄是", p.getAge())
	fmt.Println("Person set 年龄是", (&p).setAge())
	fmt.Println(p)

	// 挎包使用结构体
	w := model.CreateWroker("打工人李明")
	fmt.Println("挎包 Worker", w.GetName())
}

func printBook(book Book) {
	fmt.Printf("参数book 值为: %+v\n", book)
}

func printTeacher(t *Teacher) {
	// t是指针 t其实就是地址 应该给这个地址的指向的对象的字段赋值
	// 为了符合程序员的编程习惯 go提供了简化的赋值方式 也就是说也可以通过t.Name="小七"进行赋值。
	(*t).Name = "小六"
	fmt.Println("老师结构体内部", *t)
}