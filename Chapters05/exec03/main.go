package main

import "fmt"

func main() {

	//1.使用 switch 把小写类型的 char 型转为大写(键盘输入)。
	//只转换 a, b, c, d, e. 其它的输出“other”。

	// var char byte
	// fmt.Println("请输入你想转成大写的字母")
	// fmt.Scanf("%c", &char)

	// switch char {
	// case 'a':
	// 	fmt.Println("A")
	// case 'b':
	// 	fmt.Println("B")
	// case 'c':
	// 	fmt.Println("C")
	// case 'd':
	// 	fmt.Println("D")
	// case 'e':
	// 	fmt.Println("E")
	// default:
	// 	fmt.Println("other")
	// }

	//2.对学生成绩大于 60 分的，输出“合格”。低于 60 分的，输出“不合格”。
	//(注：输入的成绩不能大于 100)

	// 自己写的
	// var score float64
	// fmt.Println("请输入你的成绩")
	// fmt.Scanln(&score)

	// switch {
	// case score > 60 && score < 100:
	// 	fmt.Println("合格！")
	// case score < 60:
	// 	fmt.Println("不合格！")
	// default:
	// 	fmt.Println("输入成绩不能大于100分")
	// }

	// 老师写的
	// var score float64
	// fmt.Println("请输入你的成绩")
	// fmt.Scanln(&score)

	// switch int(score / 60) {
	// case 1:
	// 	fmt.Println("合格！")
	// case 0:
	// 	fmt.Println("不合格！")
	// default:
	// 	fmt.Println("输入有误")
	// }

	//3.根据用户指定月份，打印该月份所属的季节。
	// 3,4,5春季6,7,8夏季9,10,11秋季12,1,2冬季

	// var month byte
	// fmt.Println("请输入月份")
	// fmt.Scanln(&month)

	// switch month {
	// case 3, 4, 5:
	// 	fmt.Println("春季")
	// case 6, 7, 8:
	// 	fmt.Println("夏季")
	// case 9, 10, 11:
	// 	fmt.Println("秋季")
	// case 12, 1, 2:
	// 	fmt.Println("冬季")
	// default:
	// 	fmt.Println("输入有误")
	// }

	//4.根据用户输入显示对应的星期时间（string），
	//如果“星期一”，显示“干煸豆角”
	//如果“星期二”，显示“醋溜土豆”
	//如果“星期三”，显示“红烧狮子头”
	//如果“星期四”，显示“油炸花生米”
	//如果“星期五”，显示“蒜蓉扇贝”
	//如果“星期六”，显示“东北乱炖”
	//如果“星期日”，显示“大盘鸡”

	var week string
	fmt.Println("今天是星期几")
	fmt.Scanln(&week)

	switch week {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("醋溜土豆")
	case "星期三":
		fmt.Println("红烧狮子头")
	case "星期四":
		fmt.Println("油炸花生米")
	case "星期五":
		fmt.Println("蒜蓉扇贝")
	case "星期六":
		fmt.Println("东北乱炖")
	case "星期日":
		fmt.Println("大盘鸡")
	default:
		fmt.Println("无效输入")
	}

}
