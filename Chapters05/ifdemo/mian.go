package main

import "fmt"

func main() {
	//单分支
	//编写一个程序，可以输入人的年龄，如果该同志的年龄大于18岁，则输出“
	//你年龄大于18，要对自己的行为负责！”

	//分析
	//1. 年龄 ==> var age int
	//2. 从控制台接受一个输入 fmt.Scanln（&age）
	//3. if判断

	/*var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18,要对自己的行为负责！")
	} */

	// golang支持在if中，直接定义一个变量，比如下面
	// if age := 20; age > 18 {
	// 	fmt.Println("你年龄大于18,要对自己的行为负责！")
	// }

	//双分支
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18,要对自己的行为负责！")
	} else {
		fmt.Println("你的年龄不大这次放过你了")
	}
}
