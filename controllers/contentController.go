package controllers

import (
	"backend-berita/initializers"
	"backend-berita/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllContent(c *gin.Context) {
	var content []models.Content

	result := initializers.DB.Find(&content)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal memuat content"})
		return
	}

	if len(content) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ditemukan content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": content})
}

func GetContent(c *gin.Context) {
	var content models.Content

	result := initializers.DB.First(&content, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal memuat content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": content})
}

func CreateContent(c *gin.Context) {
	var content models.Content

	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&content)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membuat content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": content})
}

func UpdateContent(c *gin.Context) {
	var content models.Content

	result := initializers.DB.First(&content, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal memuat content"})
		return
	}

	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = initializers.DB.Save(&content)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengupdate content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": content})
}

func DeleteContent(c *gin.Context) {
	var content models.Content

	result := initializers.DB.First(&content, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal memuat content"})
		return
	}

	result = initializers.DB.Delete(&content)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal menghapus content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Content berhasil dihapus"})
}
