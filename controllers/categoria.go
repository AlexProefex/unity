package controllers

import (
	"net/http"
	"strconv"
	"unity/service"
	"unity/types"
	"unity/utils"

	"github.com/gin-gonic/gin"
)

func GetAllCategorias(c *gin.Context) {
	categoria, err := service.ServiceGetAllCategoria()
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func RegistrarCategoria(c *gin.Context) {
	var err error
	var input types.CategoriaRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceSaveCategoria(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoria Registrada "})
}

func ActualizarCategoria(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidID})
		return
	}
	uid := uint(id)
	var input types.CategoriaUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceUpdateCategoria(input, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoria Actualizada"})
}

func GetCategoriaById(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidID})
		return
	}
	uid := uint(id)
	categoria, err := service.ServiceGetCategoriaByID(uid)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func GetChallenge(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.InvalidID})
		return
	}
	uid := uint(id)
	categoria, err := service.ServiceGetChallenge(uid)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categoria)
}

func GenerateChallenge(c *gin.Context) {
	var input types.CategoriaChallenge
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.ServiceSetChallenge(input)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Challenge creado")
}

func GenerateMiniChallenge(c *gin.Context) {
	var input types.CategoriaMiniChallenge
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	locaciones, err := service.ServiceSetMiniChallenge(input)
	if err != nil {
		if err.Error() == utils.Not_found {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"miniChallenges": locaciones})
}

func RestartMiniChallenge(c *gin.Context) {
	var input types.CategoriaChallenge
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.ServiceRemoveMiniChallenge()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
