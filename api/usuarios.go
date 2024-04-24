package api

// Register handles user registration
/*func Register(c *gin.Context) {
var input types.UsuariosRegister

// Bind and validate input
if err := c.ShouldBindJSON(&input); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

// Create a new user
user := model.Usuarios{
	Nombre:             input.Nombre,
	Apellidos:          input.Apellidos,
	Correo_electronico: input.Correo_electronico,
	Codigo_pais:        input.Codigo_pais,
	Celular:            input.Celular,
	Genero:             input.Genero,
	Fecha_nacimiento:   input.Fecha_nacimiento,
	Pais:               input.Pais,
	Nacionalidad:       input.Nacionalidad,
	Password:           input.Password,
}

fmt.Println((user))

// Save the user to the database
/*_, err := user.SaveUser()
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}*/

//c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
//}
