package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ful/RESTful/src/storage"
)

type ViewHandler struct {
	dataStorage *storage.Storage
}

func NewViewHandler(newStorage *storage.Storage) *ViewHandler {
	return &ViewHandler{
		dataStorage: newStorage,
	}
}

func (v *ViewHandler) GetAlbums(c *gin.Context) {
	showHeader(c)

	a := (*v.dataStorage).GetAlbums()
	for _, item := range a {
		c.HTML(http.StatusOK, "view.html", gin.H{
			"ID":     item.ID,
			"Title":  item.Title,
			"Artist": item.Artist,
			"Price":  item.Price,
		})
	}
}

func (v *ViewHandler) GetAlbum(c *gin.Context) {
	showHeader(c)

	id, err := strconv.Atoi(c.Param("id"))
	if !IsIdExists(v.dataStorage, id) {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message:": "bad request"})
		return
	}
	a := (*v.dataStorage).GetAlbum(id)
	c.HTML(http.StatusOK, "view.html", gin.H{
		"ID":     a.ID,
		"Title":  a.Title,
		"Artist": a.Artist,
		"Price":  a.Price,
	})
}

func (v *ViewHandler) PostAlbum(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "post.html", gin.H{
			"ID":     0,
			"Title":  "",
			"Artist": "",
			"Price":  0.0,
		})
	case "POST":
		al, err := formAlbum(c)
		if err != nil {
			c.Redirect(http.StatusNotModified, "/albums/view")
		}
		(*v.dataStorage).CreateAlbum(al)
		c.Redirect(http.StatusSeeOther, "/albums/view")
	default:
		c.Redirect(http.StatusNotModified, "/albums/view")
	}
}

func formAlbum(c *gin.Context) (*storage.Album, error) {
	id, err := strconv.Atoi(c.Request.FormValue("fid"))
	if err != nil {
		return nil, err
	}
	title := c.Request.FormValue("ftitle")
	artist := c.Request.FormValue("fartist")
	price, err := strconv.ParseFloat(c.Request.FormValue("fprice"), 64)
	if err != nil {
		return nil, err
	}
	return &storage.Album{
		ID:     id,
		Title:  title,
		Artist: artist,
		Price:  price,
	}, nil
}

func (v *ViewHandler) DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(http.StatusNotModified, "/albums/view")
	}
	switch c.Request.Method {
	case "GET":
		a := (*v.dataStorage).GetAlbum(id)
		c.HTML(http.StatusOK, "delete.html", gin.H{
			"ID":     a.ID,
			"Title":  a.Title,
			"Artist": a.Artist,
			"Price":  a.Price,
		})
	case "POST":
		(*v.dataStorage).DeleteAlbum(id)
		c.Redirect(http.StatusSeeOther, "/albums/view")
	default:
		c.Redirect(http.StatusNotModified, "/albums/view")
	}
}

func (v *ViewHandler) UpdateAlbum(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Redirect(http.StatusNotModified, "/albums/view")
		}

		a := (*v.dataStorage).GetAlbum(id)
		c.HTML(http.StatusOK, "put.html", gin.H{
			"ID":     a.ID,
			"Title":  a.Title,
			"Artist": a.Artist,
			"Price":  a.Price,
		})
	case "POST":
		al, err := formAlbum(c)
		if err != nil {
			c.Redirect(http.StatusNotModified, "/albums/view")
		}
		(*v.dataStorage).UpdateAlbum(al.ID, al)
		c.Redirect(http.StatusSeeOther, "/albums/view")
	default:
		c.Redirect(http.StatusNotModified, "/albums/view")
	}
}

func showHeader(c *gin.Context) {
	c.HTML(http.StatusOK, "header.html", gin.H{})
}
