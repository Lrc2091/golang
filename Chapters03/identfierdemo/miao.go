package main
import (
	"fmt"
)
//演示golang标识符的使用
func main(){
	//Golang中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v", num, Num)

	//标识符不能包含空格
	//var ab c int =30

	// _是空标识符，用来占用
	// var _ int = 40 //error
	// fmt.Printl(_)

	//语法可以通过，但是要求不能使用
	var int int = 90
	fmt.Println(int)
}