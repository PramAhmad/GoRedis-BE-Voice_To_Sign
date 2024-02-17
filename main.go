package main

import (
	"backend-berita/controllers"
	"backend-berita/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	// enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Status(http.StatusOK)
		c.Next()
	})

	r.Static("/images", "./images")
	// Middleware untuk menetapkan pembatasan ukuran file
	r.Use(func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 500000)

		c.Next()
	})
	// auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.GenerateToken)
	// images
	r.GET("/image", controllers.GetImages)
	r.GET("/image/:id", controllers.GetsingleImage)
	r.GET("/images", controllers.GetAllImages)
	r.POST("/images", controllers.UploadImages)
	r.PUT("/image/:id", controllers.UpdateImage)
	r.DELETE("/image/:id", controllers.DeleteImage)
	r.Run()
	// faker
}
