package controllers

import (
	"net/http"
	"strconv"
	"unity/service"
	"unity/types"
	"unity/utils"

	"github.com/gin-gonic/gin"
)

func GetAllLocacion(c *gin.Context) {
	locacion, err := service.ServiceGetAllLocacion()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, locacion)
}

func RegistrarLocacion(c *gin.Context) {
	var err error
	var input types.LocacionRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceSaveLocacion(input)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "locacion registrada "})
}

func ActualizarLocacion(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := uint(id)
	var input types.LocacionUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceUpdateLocacion(input, uid)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "locacion actualizada "})
}

func GetLocacionById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := uint(id)
	locacion, err := service.ServiceGetLocacionByID(uid)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, locacion)
}
