package main

import "fmt"

func main() {
	// 参加百米运动会，如果用时8秒以内进入决赛
	// 否则提示淘汰，并且根据性别提示进入男子组或女子组
	// 输入成绩与性别

	// var second float64
	// var sex string
	// fmt.Println("请输入你的百米成绩：")
	// fmt.Scanln(&second)
	// fmt.Println("请输入你的性别")
	// fmt.Scanln(&sex)

	// if second < 8.0 {
	// 	if sex == "男" {
	// 		fmt.Println("恭喜你进入男子组百米决赛")
	// 	} else {
	// 		fmt.Println("恭喜你进入女子组百米决赛")
	// 	}
	// } else {
	// 	fmt.Println("很遗憾你被淘汰了..")
	// }

	// 出票系统：根据淡旺季的月份和年龄，打印票价

	// 4-10 旺季：
	// 成人（18-60）：60
	// 儿童（<18）：半价
	// 老人（>60）：1/3
	// 淡季：
	// 成人：40
	// 其他：20

	// 自己写的：
	// var month int
	// fmt.Println("请输入游玩月份：")
	// fmt.Scanln(&month)

	// if month >= 4 && month <= 10 {
	// 	var age int
	// 	fmt.Println("请输入您的年龄：")
	// 	fmt.Scanln(&age)
	// 	var price float64 = 60
	// 	if age >= 18 && age < 60 {
	// 		fmt.Println("您的票价是：", price)
	// 	} else if age < 18 {
	// 		fmt.Println("您的票价是：", price/2)
	// 	} else if age > 60 {
	// 		fmt.Println("您的票价是：", price/3)
	// 	}
	// } else {
	// 	var age int
	// 	fmt.Println("请输入您的年龄：")
	// 	fmt.Scanln(&age)
	// 	var price1 float64 = 40
	// 	if age > 18 && age < 60 {
	// 		fmt.Println("您的票价是：", price1)
	// 	} else {
	// 		fmt.Println("您的票价是：", price1/2)
	// 	}
	// }

	// 老师写的：
	var month byte
	var age byte
	var price float64 = 60.0
	fmt.Println("请输入游玩月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入您的年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("%v 月 年龄 %v 票价 %v", month, age, price/3)
		} else if age >= 18 {
			fmt.Printf("%v 月 年龄 %v 票价 %v", month, age, price)
		} else {
			fmt.Printf("%v 月 年龄 %v 票价 %v", month, age, price/2)
		}
	} else {
		//淡季
		var age int
		if age > 18 && age < 60 {
			fmt.Println("淡季成人 票价 40")
		} else {
			fmt.Println("淡季成人 票价 20")
		}
	}
}
