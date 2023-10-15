package Controllers


import (
	"learn_go/gin_crud/database"
	"learn_go/gin_crud/models"
	// "fmt"
	"github.com/gin-gonic/gin"
)
func PostBlog(c *gin.Context){
	//get data from request body
	var body struct{
		Title string `json:"title"`
		Body string `json:"body"`
	}
	c.Bind(&body)
	//create data
	post := models.Post{Title:body.Title,Body:body.Body} 
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