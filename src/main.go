package main

import (
	"ful/RESTful/src/handler"
	logger "ful/RESTful/src/logs"
	"ful/RESTful/src/storage"
)

func main() {
	dbconfig := storage.DBConfig{
		Host:     "0.0.0.0",
		Port:     5432,
		User:     "user",
		Password: "qwerty",
		DBName:   "user",
		SSLMode:  "disable",
	}

	var dataStorage storage.Storage
	var reqHandler, reqViewHandler handler.RequestHandler
	var logs logger.Logger

	logs = logger.NewLogrusLogger("./logs/logs.txt")
	logs.CreateLogger()
	defer logs.CloseLogger()

	dataStorage = storage.NewPostgresStorage()
	dataStorage.ConnectToDatabase(&dbconfig)
	defer dataStorage.CloseConnection()
	reqHandler = handler.NewDefaultHandler(&dataStorage, &logs)
	reqViewHandler = handler.NewViewHandler(&dataStorage)

	router := handler.InitRouter(&reqHandler, &reqViewHandler)

	router.Run(":8080")
}
