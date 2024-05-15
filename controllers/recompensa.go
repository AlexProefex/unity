package controllers

import (
	"net/http"
	"strconv"
	"unity/service"
	"unity/types"
	"unity/utils"

	"github.com/gin-gonic/gin"
)

func GetAllRecompensa(c *gin.Context) {
	recompensa, err := service.ServiceGetAllRecompensa()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recompensa)
}

func RegistrarRecompensa(c *gin.Context) {
	var err error
	var input types.RecompensaRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceSaveRecompensa(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "recompensa registrada "})
}

func ActualizarRecompensa(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid := uint(id)

	var input types.RecompensaUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceUpdateRecompensa(input, uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "recompensa actualizada "})
}

func GetRecompensaById(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid := uint(id)

	recompensa, err := service.ServiceGetRecompensaByID(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recompensa)
}

func GetRecompensaByUserId(c *gin.Context) {

	recompensa, err := service.ServiceGetRecompensaByID(c.MustGet("userID").(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recompensa)
}

func CanjearRecompesasInsignia(c *gin.Context) {

	id := c.MustGet("userID").(uint)

	cantidad := c.MustGet("cantidad").(int)

	recompensa, err := service.ServiceCobrarAgregarRecompensaInsignia(id, cantidad)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recompensa)
}

func CanjearRecompesasPuntos(c *gin.Context) {

	id := c.MustGet("userID").(uint)
	puntos := c.MustGet("puntos").(int)

	recompensa, err := service.ServiceCobrarAgregarRecompensaPuntos(id, puntos)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recompensa)
}

func GenerateQRToken(c *gin.Context) {
	var input types.UsuariosValidateCP
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := service.ServiceGenerarPagoCodigoQR(input, c.MustGet("userID").(uint))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	utils.TokenCookie(c)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
