package routes

import (
	"unity/initialize"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// albums slice to seed record album data.

// getAlbums responds with the list of all albums as JSON.
/*func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, api.ShowAlbum())
}*/

func Routes() {

	router := gin.Default()
	initialize.ConnectionDB()

	url := ginSwagger.URL("http://localhost:8080/swagger/docs.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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

	/*api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// Register and login routes
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)

		}

		//api.GET("/albums", controllers.ShowAlbum)
		// Profile route with JWT authentication middleware
		//api.GET("/profile", middleware.JwtAuthMiddleware(), controllers.Profile)
		v1 := api.Group("/v1")
		v1.Use(middleware.JwtAuth())
		api.GET("/albums", controllers.ShowAlbum)
		v1.GET("/all", controllers.GetAllUsers)
		v1.GET("/hello", controllers.ShowAlbum)
		v1.GET("/user", controllers.UserData)

	}*/

	router.Run("localhost:8080")

}

//r := gin.Default()

// Load database connection
//models.ConnectDataBase()

// Group routes

// Run the server
//r.Run(":8080")
