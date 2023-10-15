package Controllers


import (
	"learn_go/gin_crud/database"
	"learn_go/gin_crud/models"
	// "fmt"
	"github.com/gin-gonic/gin"
)
func PostBlog(c *gin.Context){
	//get data from request body
	//create data
	post := models.Post{Title:"Telugu Heros",Body:"Chiru,Nag,Venky,Bal"} 
	result := database.DB.Create(&post)
	if result.Error != nil{
		c.Status(400)
		return
	}
	c.JSON(200,gin.H{
		"title":post.Title,
	})
	// fmt.Println("ppppppppppp")

}