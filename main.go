package main
import (
	"learn_go/gin_crud/database"
	"learn_go/gin_crud/initializers"
	"github.com/gin-gonic/gin"
    "fmt"
)

func main(){
	initializers.Loadenv()
	// DATABASE.ConnectToDb()
	database.ConnectToDb()
	//create router
	//for testing from browser
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.JSON(200,"Hello Shailesh and Hello World")
	})
	fmt.Println("Hello World")
	r.Run() //default port 8080

}
