package controllers

import (
	"backend-berita/auth"
	"backend-berita/initializers"
	"backend-berita/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type TokenReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenReq
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": "Masukan Usernamae dan Password "})
		c.Abort()
		return
	}
	// cek apakah user ada di database
	var user models.User
	result := initializers.DB.Where("username = ?", request.Username).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "User tidak ditemukan"})
		c.Abort()
		return
	}
	// cek apakah passwordnya cocok implementasi gunakan go routine

	err := ComparePassword(user.Password, request.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Password salah"})
		c.Abort()
		return
	}

	// // jika ada, generate token go routine get token
	tokenString, err := auth.GenerateJwt(user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error saat membuat token"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"token": tokenString})

}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
