package main
// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

func main() {

	//int ，uint ，rune ， byte的使用
	// var a int = 99999
	// var b uint = 88
	// var c byte = 255
	// fmt.Println("a=", a,"b=", b, "c=", c)

	//整型的使用细节
	var n1 = 100 // ? n1 是什么类型
	//这里我们给介绍一下如何查看某个变量的数据类型
	//fmt.Printf() 可以用于做格式化输出。
	fmt.Printf("n1 的 类型 %T \n", n1)

	//如何在程序查看某个变量的 占用字节大小和数据类型 (使用较多)
	var n2 int64 = 10
	//unsafe.sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的类型 %T n2占用的字节数是 %d \n", n2, unsafe.Sizeof(n2))

	//Golang程序中整型变量在使用时，遵守保小不保大的原则，
	//即:在保证程序正确运行下，尽量使用占用空间小的数据类型
	var age byte = 90
	fmt.Printf("age 的 类型 %T \n", age)
}
	
