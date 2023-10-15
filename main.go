package main
import (
	"learn_go/gin_crud/Controllers"
	"learn_go/gin_crud/models"
	"learn_go/gin_crud/database"
	"learn_go/gin_crud/initializers"
	"github.com/gin-gonic/gin"
    "fmt"
)
var post  = []models.Post{}

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
	r.POST("/post",Controllers.PostBlog)
	fmt.Println("Hello World")
	r.Run() //default port 8080

}
