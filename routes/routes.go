package routes

import (
	"unity/controllers"
	"unity/initialize"
	"unity/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() {
	router := gin.New()
	//router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	initialize.ConnectionDB()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
	//handle error
	}*/

	/*authorized := router.Group("/home", gin.BasicAuth(gin.Accounts{
		"user1": "love",
		"user2": "god",
		"user3": "sex",

	}))

	authorized.GET("/secret", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"success": true})

		})*/
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// Register and login routes
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}
		internalServer := api.Group("internal")
		internalServer.POST("restart-challenge", controllers.RestartMiniChallenge)
		//api.GET("/albums", controllers.ShowAlbum)
		// Profile route with JWT authentication middleware
		//api.GET("/profile", middleware.JwtAuthMiddleware(), controllers.Profile)
		customer := api.Group("customer")
		{
			customer.Use(middleware.JwtQRAuth())
			customer.POST("/puntos", controllers.CanjearRecompesasPuntos)
			customer.POST("/insignias", controllers.CanjearRecompesasInsignia)

		}
		recover := api.Group("recover")
		{
			recover.POST("/", controllers.RecuperarContrasena)
		}

		v1 := api.Group("/v1")
		v1.Use(middleware.JwtAuth())
		categoria := v1.Group("/categoria")
		{
			categoria.GET("/", controllers.GetAllCategorias)
			categoria.POST("/challenge", controllers.GenerateChallenge)
			categoria.POST("/mini-challenge", controllers.GenerateMiniChallenge)
			//categoria.POST("/restart-challenge", controllers.RestartMiniChallenge)

			categoria.GET("/routes/:id", controllers.GetChallenge)

			categoria.POST("/", controllers.RegistrarCategoria)
			categoria.PUT("/:id", controllers.ActualizarCategoria)
			categoria.GET("/:id", controllers.GetCategoriaById)

		}
		locacion := v1.Group("/locacion")
		{
			locacion.GET("/", controllers.GetAllLocacion)
			locacion.POST("/", controllers.RegistrarLocacion)
			locacion.PUT("/:id", controllers.ActualizarLocacion)
			locacion.GET("/:id", controllers.GetCategoriaById)

		}
		recompensa := v1.Group("/recompensa")
		{
			recompensa.GET("/current", controllers.GetRecompensaByUserId)
			recompensa.GET("/:id", controllers.GetRecompensaById)
			recompensa.GET("/", controllers.GetAllRecompensa)
			recompensa.POST("/", controllers.RegistrarRecompensa)
			recompensa.PUT("/:id", controllers.ActualizarRecompensa)
			recompensa.POST("/gift", controllers.CanjearRecompesasInsignia)
			recompensa.POST("/points", controllers.CanjearRecompesasPuntos)

		}
		premio := v1.Group("/premio")
		{
			premio.GET("/descuento", controllers.GetAllPremioDescuento)
			premio.GET("/regalo", controllers.GetAllPremioRegalo)
			premio.POST("/qr", controllers.GenerateQRToken)
			premio.GET("/", controllers.GetAllPremio)
			premio.POST("/", controllers.RegistrarPremio)
			premio.PUT("/:id", controllers.ActualizarPremio)
			premio.GET("/:id", controllers.GetPremioById)

		}
		usuario := v1.Group("/usuario")
		{

			usuario.GET("/list", controllers.GetAllUsers)
			usuario.GET("/profile", controllers.GetUserById)
			usuario.POST("/change-password", controllers.CambiarContrasena)
			usuario.POST("/update-perfil", controllers.UpdatePerfil)
			usuario.POST("/add-points", controllers.AsignarPuntos)
		}
	}
	router.Run("localhost:8080")
}

//r := gin.Default()

// Load database connection
//models.ConnectDataBase()

// Group routes

// Run the server
//r.Run(":8080")
