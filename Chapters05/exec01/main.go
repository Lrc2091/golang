package main

import "fmt"

func main() {
	//岳小鹏参加golang考试，她和父亲达成承诺：
	//如果：
	//成绩为100分，奖励一台BMW
	//成绩为（80，99]时，奖励一台iphone15 pro max
	//当成绩为[60，80]，奖励一个ipad
	//其他时什么奖励都没有
	//请从键盘输入岳小鹏的期末成绩，并加以判断

	// 分析思路
	// 1. score 分数变量 int
	// 2. 选择多分支流程控制
	// 3. 成绩从键盘输入 fmt.Scanln

	// var score int
	// fmt.Println("请输入岳小鹏的期末成绩：")
	// fmt.Scanln(&score)
	// if score == 100 {
	// 	fmt.Println("奖励一辆BMW")
	// } else if score > 80 && score <= 99 {
	// 	fmt.Println("奖励一台iphone15 pro max")
	// } else if score >= 60 && score <= 80 {
	// 	fmt.Println("奖励一台ipad")
	// } else {
	// 	fmt.Println("什么都不奖励")
	// }

	// // 使用陷阱  只会输出 ok1

	// var n int = 10
	// if n > 9 {
	// 	fmt.Println("ok1") //输出
	// } else if n > 6 {
	// 	fmt.Println("ok2")
	// } else if n > 3 {
	// 	fmt.Println("ok3")
	// } else {
	// 	fmt.Println("ok4")
	// }

	// 案例演示3
	// 求ax2+bx+c=0方程的根。a,b,c分别为函数的参数，如果：b2-4ac>0，则有两个解；
	// b2-4ac=0，则有一个解；b2-4ac<O，则无解；
	// 提示1：x1=（-b+sqrt（b2-4ac））/2a X2=(-b-sqrt(b2-4ac))/2a
	// 提示2： math.Sgrt（num）；可以求平方根 需要引入 math包
	// 测试数据：3,100,6

	// 分析思路
	// 1, a,b,c，是三个float64
	// 2，使用到给出的数学公式
	// 3，使用到多分支
	// 4，使用math.Squr 方法 => 手册

	// var a float64 = 2.0
	// var b float64 = 4.0
	// var c float64 = 2.0

	// y := b*b - 4*a*c
	// 多分支判断
	// if y > 0 {
	// 	x1 := (-b + math.Sqrt(y)/2*a)
	// 	x2 := (-b - math.Sqrt(y)/2*a)
	// 	fmt.Printf("x1=%v x2=%v", x1, x2)
	// } else if y == 0 {
	// 	x := (-b + math.Sqrt(y)/2*a)
	// 	fmt.Printf("x=%v", x)
	// } else {
	// 	fmt.Println("无解")
	// }

	// 大家都知道，男大当婚，女大当嫁。那么女方家长要嫁女儿，当然要提出一定的条件，
	// 高：180cm以上；富：财富1千万以上，帅：是。 条件从控制台输人。
	//1）如果这三个条件同时满足，则，“我一定要嫁给他！！！”
	//2）如果三个条件有为真的情况，则：“嫁吧，比上不足，比下有余。”
	//3）如果三个条件都不满足，则，“不嫁！”

	var height int
	var money float32
	var featureas bool
	fmt.Println("请输入身高(cm)")
	fmt.Scanln(&height)
	fmt.Println("请输入财富(千万)")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅(true/false)")
	fmt.Scanln(&featureas)

	if height > 180 && money >= 1.0 && featureas {
		fmt.Println("我一定要嫁给他！！！")
	} else if height > 180 || money >= 1.0 || featureas {
		fmt.Println("嫁吧，比上不足，比下有余")
	} else {
		fmt.Println("不嫁...")
	}

}
