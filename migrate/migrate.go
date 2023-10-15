package main

import (
	"learn_go/gin_crud/database"
	"learn_go/gin_crud/initializers"
	"learn_go/gin_crud/models"
)

func init(){
	initializers.Loadenv()
	database.ConnectToDb()
	
}

func main(){
	database.DB.AutoMigrate(&models.Post{})
}
