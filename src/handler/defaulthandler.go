package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	logger "ful/RESTful/src/logs"
	"ful/RESTful/src/storage"
)

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
		(*d.logHandler).WriteError(fmt.Sprintf("GET: unknown album id: %d", id), c.RemoteIP())
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		(*d.logHandler).WriteError(fmt.Sprintf("GET: incorrect input id: %d", id), c.RemoteIP())
		return
	}
	c.IndentedJSON(http.StatusOK, (*d.dataStorage).GetAlbum(id))
}

func (d *DefaultHandler)PostAlbum(c *gin.Context) {
	newAlbum := storage.Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		(*d.logHandler).WriteError(fmt.Sprintf("POST: failed to bind json"), c.RemoteIP())
		return
	}
	if IsIdExists(d.dataStorage, newAlbum.ID) {
		c.IndentedJSON(http.StatusNotImplemented, nil)
		(*d.logHandler).WriteError(fmt.Sprintf("POST: id already exists: %d", newAlbum.ID), c.RemoteIP())
		return
	}
	(*d.dataStorage).CreateAlbum(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (d *DefaultHandler)DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		(*d.logHandler).WriteError(fmt.Sprintf("DELETE: incorrect input id: %d", id), c.RemoteIP())
		return
	}
	if !IsIdExists(d.dataStorage, id) {
		c.IndentedJSON(http.StatusNotFound, nil)
		(*d.logHandler).WriteError(fmt.Sprintf("DELETE: non-existent id: %d", id), c.RemoteIP())
		return
	}
	(*d.dataStorage).DeleteAlbum(id)
	c.IndentedJSON(http.StatusNoContent, nil)
}

func (d *DefaultHandler)UpdateAlbum(c *gin.Context) {
	newAlbum := storage.Album{}
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		(*d.logHandler).WriteError(fmt.Sprintf("UPDATE: failed to bind json"), c.RemoteIP())
		return
	}
	correctID, err := strconv.Atoi(c.Request.URL.Path[len("/albums/"):])
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		(*d.logHandler).WriteError(fmt.Sprintf("UPDATE: failed while getting id value"), c.RemoteIP())
		return
	}
	if !IsIdExists(d.dataStorage, correctID) {
		c.IndentedJSON(http.StatusNotFound, nil)
		(*d.logHandler).WriteError(fmt.Sprintf("UPDATE: non-existent id: %d", correctID), c.RemoteIP())
		return
	}
	newAlbum.ID = correctID
	(*d.dataStorage).UpdateAlbum(newAlbum.ID, &newAlbum)
	c.IndentedJSON(http.StatusCreated, (*d.dataStorage).GetAlbum(newAlbum.ID))
}