package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {	
	c.IndentedJSON(http.StatusOK, dataStorage.GetAlbums())
}

func postAlbum(c *gin.Context) {
	newAlbum := Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	dataStorage.CreateAlbum(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
	}
	c.IndentedJSON(http.StatusOK, dataStorage.GetAlbum(id))
}

func deleteAlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
	}
	dataStorage.DeleteAlbum(id)
	c.IndentedJSON(http.StatusNoContent, nil)
}

func updateAlbumById(c *gin.Context) {
	newAlbum := Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	dataStorage.UpdateAlbum(newAlbum.ID, &newAlbum)
	c.IndentedJSON(http.StatusCreated, dataStorage.GetAlbum(newAlbum.ID))
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/albums", postAlbum)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE("/albums/:id", deleteAlbumById)
	router.PUT("/albums/:id", updateAlbumById)
	return router
}