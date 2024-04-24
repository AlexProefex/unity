package controllers

import (
	"net/http"
	"unity/service"

	"github.com/gin-gonic/gin"
)

/*
func ShowAlbum() []types.Album {
	return service.GetDataAlbums()
}*/

func ShowAlbum(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, service.GetDataAlbums())
}

/*
c.IndentedJSON(http.StatusOK, api.ShowAlbum())
func getAlbums(c *gin.Context) {

}*/
