package main

import "fmt"

func test(char byte) byte {
	return char + 1
}
func main() {
	// 案例
	// 请编写一个程序，该程序可以接收一个字符，比如: a,b,c,d,e,f,g
	// a 表示星期一，b 表示星期二 …	根据用户的输入显示相依的信息.
	// 要求使用 switch  语句完成

	// 分析思路
	// 1.定义一个变量接受字符
	// 2.使用switch完成
	// var key byte
	// fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	// fmt.Scanf("%c", &key)

	// switch test(key) + 1 {
	// case 'a':
	// 	fmt.Println("周一，猴子穿新衣")
	// case 'b':
	// 	fmt.Println("周二，猴子当小二")
	// case 'c':
	// 	fmt.Println("周三，猴子爬雪山")
	// case 'd':
	// 	fmt.Println("周四，猴子穿新鞋")
	// case 'e':
	// 	fmt.Println("周五，猴子舞狮子")
	// case 'f':
	// 	fmt.Println("周六，猴子溜溜梅")
	// case 'g':
	// 	fmt.Println("周日，猴子躺板板")
	// default:
	// 	fmt.Println("输入有误。")
	// }

	// var n1 int32 = 20
	// // var n2 int64 = 20
	// var n3 int32 = 20
	// switch n1 {
	// //case n2: // 错误，原因是 n2 的数据类型与n1不一致
	// //	fmt.Println("ok1")
	// case n3, 10, 5:
	// 	fmt.Println("ok2") // case 后面可以有多个表达式
	// //case 5 : 因为前面有常量5了，重复就会报错
	// //	fmt.Println("ok3")
	// default: // default 不是必须的 可有可无
	// 	fmt.Println("没有匹配到")
	// }

	//switch 后也可以不带表达式，类似 if --else 分支来使用。【案例演示】
	var age int = 10

	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}

	// case中也可以对age 范围进行判断
	var score int = 90
	switch {
	case score > 90:
		fmt.Println("成绩优秀！")
	case score >= 70:
		fmt.Println("成绩优良！")
	case score >= 60:
		fmt.Println("成绩及格.")
	default:
		fmt.Println("不及格")
	}

	// switch 后可以直接声明/定义一个变量，分号结束，不推荐
	switch graed := 90; { //在golang中，可以这样写
	case graed > 90:
		fmt.Println("成绩优秀!~")
	case graed >= 70:
		fmt.Println("成绩优良!~")
	case graed >= 60:
		fmt.Println("成绩及格.~")
	default:
		fmt.Println("不及格")
	}

	//switch 穿透 fallthrough，
	//如果在 case 语句块后增加 fallthrough
	//则会继续执行下一个 case，也叫 switch 穿透
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough // 默认只能穿透一层
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

	//Type Switch：switch 语句还可以被
	//用于 type-switch 来判断某个 interface 变量中实际指向的变量类型
	//【还没有学 interface,  先体验一把】
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型~ :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("没有匹配到")
	}
}
