package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestHandler interface {
	GetAlbums(c *gin.Context)
	GetAlbum(c *gin.Context)
	PostAlbum(c *gin.Context)
	DeleteAlbum(c *gin.Context)
	UpdateAlbum(c *gin.Context)
}

type DefaultHandler struct {
	dataStorage *Storage
}

func NewDefaultHandler(newStorage *Storage) *DefaultHandler {
	return &DefaultHandler{
		dataStorage: newStorage,
	}
}

func (d *DefaultHandler)GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, (*d.dataStorage).GetAlbums())
}

func (d *DefaultHandler)GetAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
	}
	c.IndentedJSON(http.StatusOK, (*d.dataStorage).GetAlbum(id))
}

func (d *DefaultHandler)PostAlbum(c *gin.Context) {
	newAlbum := Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	(*d.dataStorage).CreateAlbum(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (d *DefaultHandler)DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
	}
	(*d.dataStorage).DeleteAlbum(id)
	c.IndentedJSON(http.StatusNoContent, nil)
}

func (d *DefaultHandler)UpdateAlbum(c *gin.Context) {
	newAlbum := Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	(*d.dataStorage).UpdateAlbum(newAlbum.ID, &newAlbum)
	c.IndentedJSON(http.StatusCreated, (*d.dataStorage).GetAlbum(newAlbum.ID))
}


func InitRouter(handler *RequestHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/albums", (*handler).PostAlbum)
	router.GET("/albums", (*handler).GetAlbums)
	router.GET("/albums/:id", (*handler).GetAlbum)
	router.DELETE("/albums/:id", (*handler).DeleteAlbum)
	router.PUT("/albums/:id", (*handler).UpdateAlbum)
	return router
}