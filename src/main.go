package main

import(
	"ful/RESTful/src/storage"
	"ful/RESTful/src/handler"
)

func main() {
	dbconfig := storage.DBConfig{
		Host: "0.0.0.0",
		Port: 5432,
		User: "user",
		Password: "qwerty",
		DBName: "user",
		SSLMode: "disable",
	}

	var dataStorage storage.Storage
	var reqHandler handler.RequestHandler
	
	dataStorage = storage.NewPostgresStorage()
	dataStorage.ConnectToDatabase(&dbconfig)
	defer dataStorage.CloseConnection()
	reqHandler = handler.NewDefaultHandler(&dataStorage)
	
	router := handler.InitRouter(&reqHandler)
	//gin.SetMode("release")
	router.Run(":8080")
}