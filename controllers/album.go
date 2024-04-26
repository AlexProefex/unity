package controllers

import (
	"net/http"
	"unity/service"

	"github.com/gin-gonic/gin"
)

func ShowAlbum(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, service.GetDataAlbums())
}
