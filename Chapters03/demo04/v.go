package main
import (
	"fmt"
	//为了使用utils.go文件的变量或者函数，我们需要先引入该包
	"github.com/Lrc2091/golang/demo04/model"

	"github.com/gin-gonic/gin"

)
//变量使用的注意事项
func main() {
	fmt.Println("aaa")
	//该区域的数据可以在同一类型范围内不断变换
	i  := 10
	i = 30
	i = 50
	fmt.Println("i=", i)
	// i = 1.2
	//int ，原因是不能改变数据类型

	//变量在同一个作用域(在一个函数或者在代码块)内不能重名
	// var i int = 59
	// i := 99

	//我们使用utils.go 的 herName 包名。标识符
	fmt.Println(model.HerName)


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong111",
		})
	})
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
