package main
import (
	"github.com/gin-gonic/gin"
    "fmt"
)

func main(){
	//create router
	//for testing from browser
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.JSON(200,"Hello Shailesh and Hello World")
	})

	fmt.Println("Hello World")
	r.Run() //default port 8080

}
