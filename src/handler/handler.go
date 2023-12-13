package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	logger "ful/RESTful/src/logs"
	"ful/RESTful/src/storage"
)

type RequestHandler interface {
	GetAlbums(c *gin.Context)
	GetAlbum(c *gin.Context)
	PostAlbum(c *gin.Context)
	DeleteAlbum(c *gin.Context)
	UpdateAlbum(c *gin.Context)
}

type DefaultHandler struct {
	dataStorage *storage.Storage
	logHandler *logger.Logger
}

func NewDefaultHandler(newStorage *storage.Storage, newLogHandler *logger.Logger) *DefaultHandler {
	return &DefaultHandler{
		dataStorage: newStorage,
		logHandler: newLogHandler,
	}
}

func (d *DefaultHandler)GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, (*d.dataStorage).GetAlbums())
}

func (d *DefaultHandler)GetAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if !IsIdExists(d.dataStorage, id) {
		c.IndentedJSON(http.StatusNotFound, nil)
		(*d.logHandler).WriteError(fmt.Sprintf("error: unknown album id: %d", id))
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	c.IndentedJSON(http.StatusOK, (*d.dataStorage).GetAlbum(id))
}

func (d *DefaultHandler)PostAlbum(c *gin.Context) {
	newAlbum := storage.Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	if IsIdExists(d.dataStorage, newAlbum.ID) {
		c.IndentedJSON(http.StatusNotImplemented, nil)
		return
	}
	(*d.dataStorage).CreateAlbum(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (d *DefaultHandler)DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	if !IsIdExists(d.dataStorage, id) {
		c.IndentedJSON(http.StatusNotImplemented, nil)
		return
	}
	(*d.dataStorage).DeleteAlbum(id)
	c.IndentedJSON(http.StatusNoContent, nil)
}

func (d *DefaultHandler)UpdateAlbum(c *gin.Context) {
	newAlbum := storage.Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	correctID, err := strconv.Atoi(c.Request.URL.Path[len("/albums/"):])
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	if !IsIdExists(d.dataStorage, correctID) {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}
	newAlbum.ID = correctID
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