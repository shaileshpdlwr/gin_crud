package database

 
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
  )
  var DB *gorm.DB
  var err error

  func ConnectToDb(){
  dsn := "host=localhost user=shailesh password=shailesh dbname=splitapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err!=nil{
	log.Fatal("Failed to connect to DATABASE")
  }
  }