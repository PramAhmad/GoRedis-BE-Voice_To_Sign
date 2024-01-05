package main

import (
	"backend-berita/controllers"
	"backend-berita/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	r.Static("/images", "./images")
	//images
	r.GET("/image", controllers.GetImages)
	r.POST("/image", controllers.UploadImages)
	r.Run()
}
