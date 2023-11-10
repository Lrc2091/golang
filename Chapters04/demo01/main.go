package main
import (
	"fmt"
)
//演示golang中小数类型使用
func main(){

	//重点讲解 /，%
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// 演示 % 的使用
	// 公式 a % b = a - a / b * b
	fmt.Println("10%3=",10 % 3) // 1
	fmt.Println("-10%3=",-10 % 3) // -1
	fmt.Println("10%-3=",10 % -3) // 1
	fmt.Println("-10%-3=",-10 % -3) // -1

	// ++ 和 -- 的使用
	var i int = 10
	i++ // 等价 i = i + 1
	fmt.Println("i=", i)
	i-- // 等价 i = i - 1
	fmt.Println("i=", i)
}