package main
// import "fmt"
// import "unsafe"
import (
	"fmt"
)
//演示golang中字符类型使用
func main(){

	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//当我们直接输入byte值，就是输出了的相对应的字符的码值
	// 'a' ==> 97 
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2 )

	//var c3 byte = '北' // overflow 溢出
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应码值=%d\n", c3, c3)

	var c4 int = 120 // 22269 -> '国' 120 -> 'x'
	fmt.Printf("c4=%c\n", c4)

	//字符类型是可以进行运算的，相当于一个整数，运算时是按照码值运行
	var n1 = 10 + 'a' // 10 + 97 = 107
	fmt.Println("n1=",n1)
}
	
