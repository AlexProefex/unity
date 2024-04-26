package controllers

import (
	"net/http"
	"unity/service"
	"unity/types"
	"unity/utils"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Registrar Usuario
// @Description Guardar nuevo usuario
// @Param user body types.UsuariosRegister true "Create User"
// @Produce application/json
// @Tags Usuario
// @Success 200
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var err error
	var input types.UsuariosRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = service.ServiceRegister(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser godoc
// @Summary Login Usuario
// @Description Autentificar usuario
// @Param user body types.UsuariosLogin true "Autentificar Usuario"
// @Produce application/json
// @Tags Usuario
// @Success 200
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var input types.UsuariosLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := service.ServiceLogin(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	utils.TokenCookie(c)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetUserById(c *gin.Context) {
	user, err := service.ServiceGetUserByID(c.MustGet("userID").(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ListUser godoc
// @Summary Listado de Usuarios
// @Description Listado de usuarios
// @Schemes
// @Produce application/json
// @Tags Usuario4
// @Success 200
// @Router /v1/list [get]
// @Security Bearer
func GetAllUsers(c *gin.Context) {
	user, err := service.ServiceGetAllUser()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
