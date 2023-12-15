// The package is responsible
// for routing requests and logging errors.

package handler

import (
	"github.com/gin-gonic/gin"
)

type RequestHandler interface {
	GetAlbums(c *gin.Context)
	GetAlbum(c *gin.Context)
	PostAlbum(c *gin.Context)
	DeleteAlbum(c *gin.Context)
	UpdateAlbum(c *gin.Context)
}

func InitRouter(jsonhandler *RequestHandler, htmlhandler *RequestHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/albums", (*jsonhandler).PostAlbum)
	router.GET("/albums", (*jsonhandler).GetAlbums)
	router.GET("/albums/:id", (*jsonhandler).GetAlbum)
	router.DELETE("/albums/:id", (*jsonhandler).DeleteAlbum)
	router.PUT("/albums/:id", (*jsonhandler).UpdateAlbum)

	router.GET("albums/view", (*htmlhandler).GetAlbums)
	router.GET("albums/view/:id", (*htmlhandler).GetAlbum)
	router.GET("albums/edit", (*htmlhandler).PostAlbum)
	router.POST("albums/edit", (*htmlhandler).PostAlbum)
	router.GET("albums/edit/:id", (*htmlhandler).UpdateAlbum)
	router.POST("albums/edit/:id", (*htmlhandler).UpdateAlbum)

	router.LoadHTMLFiles("./html/view.html", "./html/header.html", "./html/post.html", "./html/put.html")

	return router
}
