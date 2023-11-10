package main
import (
	"fmt"
	"strconv"
)
//演示golang中string转成基本数据类型
func main(){

	var str string = "ture"
	var b bool
	//b, _ = strconv.ParseBool(str)
	//说明
	//1.strconv.ParseBool(str) 函数会返回两个值（value bool, err error）
	//2.因为我只想获取到value bool，不想获取 err 所以我使用_忽略
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "12345678"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T b=%v\n", n1, n1)
	fmt.Printf("n2 type %T b=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T b=%v\n", f1, f1)

	var num1 string = "123123"
	var num2 int64
	num2, _ = strconv.ParseInt(num1, 10, 64)
	fmt.Printf("num2 type %T b=%v\n", num2, num2)
	var num3 int32
	num3 = int32(num2)
	fmt.Printf("num3 type %T b=%v\n", num3, num3)


	//注意：
	var str4 string = "hello"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T b=%v\n", n3, n3)

}