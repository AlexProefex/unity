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

// GetUserId godoc
// @Summary Obtener Usuario por Id
// @Description Obtener los datos del usuario por su id
// @Produce application/json
// @Tags Usuario
// @Success 200 {object}  types.UsuariosModel
// @Router /v1/user [get]
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
// @Tags Usuario
// @Success 200  {array}   types.UsuariosModel
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

func RecuperarContrasena(c *gin.Context) {

	var input types.UsuariosPassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario, err := service.ServiceRecuperarContrasena(input)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func CambiarContrasena(c *gin.Context) {
	var input types.ConfirmPassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuario, err := service.ServiceCambiarContrasena(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuario)
}
