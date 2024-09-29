package main

import (
	"backend-berita/initializers"
	"backend-berita/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Images{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Content{})
}
