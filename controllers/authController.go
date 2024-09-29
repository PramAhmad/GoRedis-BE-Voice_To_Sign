package controllers

import (
	"backend-berita/initializers"
	"backend-berita/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}


	hashPw, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error saat menghash password"})
		c.Abort()
		return
	}
	// create user
	user := models.User{
		Username: body.Username,
		Password: string(hashPw),
	}
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error saat membuat user"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "User berhasil dibuat"})

}

// func Auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token == "" {
// 			c.JSON(401, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}
// 		claims, err := auth.VerifyJwt(token)
// 		if err != nil {
// 			c.JSON(401, gin.H{"error": "Unauthorized"})
// 			c.Abort()
// 			return
// 		}
// 		c.Set("username", claims.Username)
// 		c.Next()
// 	}
// }
