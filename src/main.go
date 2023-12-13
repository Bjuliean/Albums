package main

import (
	//"github.com/gin-gonic/gin"
)

func main() {
	dbconfig := DBConfig{
		Host: "0.0.0.0",
		Port: 5432,
		User: "user",
		Password: "qwerty",
		DBName: "user",
		SSLMode: "disable",
	}

	var dataStorage Storage
	var handler RequestHandler
	
	dataStorage = NewPostgresStorage()
	dataStorage.ConnectToDatabase(&dbconfig)
	defer dataStorage.CloseConnection()
	handler = NewDefaultHandler(&dataStorage)
	router := InitRouter(&handler)
	//gin.SetMode("release")
	router.Run(":8080")
}