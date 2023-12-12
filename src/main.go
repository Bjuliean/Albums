package main

import (
	"github.com/gin-gonic/gin"
)

var dataStorage Storage

func main() {
	// dbconfig := DBConfig{
	// 	Addr: "0.0.0.0",
	// 	Port: 5432,
	// 	User: "user",
	// 	Password: "qwerty",
	// 	DB: "user",
	// }
	// dataStorage, err := NewPostgresStorage(&dbconfig)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	dataStorage = NewPostgresStorage()
	dataStorage.ConnectToDatabase()
	defer dataStorage.CloseConnection()
	router := InitRouter()
	gin.SetMode("release")
	router.Run(":8080")
}