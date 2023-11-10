package main

import "fmt"

func main() {

	//单分支和双分支练习题
	var a int = 4
	var b int = 1
	if a > 2 {
		if b > 2 {
			fmt.Println(a + b)
		}
		fmt.Println("shuaibi")
	} else {
		fmt.Println("a is =", a)
	}

	var x int = 4
	if x > 2 {
		fmt.Println("ok")
	} else {
		fmt.Println("hello")
	}

	//编写程序，声明2个int32型变量并赋值，判断两数之和，如果大于等于50，打印“hello world！”
	var q int32 = 32
	var w int32 = 19
	if q+w > 50 {
		fmt.Println("hello world!")
	}

	//编写程序，声明两个float64型变量并赋值，判断第一个数大于10.0，且第二个数小于20.0，打印两数之和
	var n1 float64 = 19.99
	var n2 float64 = 15.9
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println("和=", n1+n2)
	}

	// 定义两个变量int32，判断二者的和，是否能被3又能被5整除，打印提示信息
	// var n3 int32
	// fmt.Println("请输入第一个数字：")
	// fmt.Scanln(&n3)
	// var n4 int32
	// fmt.Println("请输入第二个数字：")
	// fmt.Scanln(&n4)
	// if (n3+n4)%3 == 0 && (n3+n4)%5 == 0 {
	// 	fmt.Println("两数之和能被3和5整除")
	// }

	//判断一个年份是否是闰年，闰年的条件是符合下面二者之一：1，年份能被4整除，但不能被100整除；
	//2，能被400整除
	var year int
	fmt.Println("请输入年份")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是闰年~", year)
	}

}
