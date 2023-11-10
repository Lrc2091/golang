package main
import(
	"fmt" 
)
func main(){

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10 //这里修改时，会到num的值变化
	fmt.Println("num = ", num)
}