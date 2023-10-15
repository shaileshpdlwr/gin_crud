package main
import (
	"github.com/gin-gonic/gin"
    "fmt"
	"log"
    "github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

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
