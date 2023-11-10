package main
import (
	"fmt"
)
func main(){

	//在golang中， ++ 和 -- 只能独立使用
	// var i int = 8
	// var j int 
	// a = i++  // 错误，i++只能独立使用
	// a = i--  // 错误，i--只能独立使用

	// if i++ > 0 {
	// 	fmt.Println("ok")
	// } // 错误

	var i int = 1
	i++
	++i //错误 在golang中没有前++和前--
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)
}