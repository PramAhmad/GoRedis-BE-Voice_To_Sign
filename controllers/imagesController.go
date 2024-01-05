package controllers

import (
	"backend-berita/initializers"
	"backend-berita/models"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"strings"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	combinedNames := c.Query("names")

	names := strings.Fields(combinedNames)

	var images []models.Images

	cacheResult, cacheErr := initializers.RedisClient().Get(c, "images").Result()
	if cacheErr == nil {
		// Data ketemmu di redis
		if err := json.Unmarshal([]byte(cacheResult), &images); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": " unmarshaling cached data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": images})
		return
	}

	result := initializers.DB.Where("name IN ?", names).Find(&images)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal memuat gambar"})
		return
	}

	if len(images) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ditemukan gambar dengan nama yang cocok"})
		return
	}

	// simpan ke redis
	jsonData, err := json.Marshal(images)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling data"})
		return
	}
	if err := initializers.RedisClient().Set(c, "images", jsonData, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error caching data dari redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": images})
}

func UploadImages(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// validate jika name kosong
	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama tidak boleh kosong"})
		return
	}

	uploadedImage, err := c.FormFile("path")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar tidak ditemukan"})
		return
	}
	// custom name file di ubah jadi dari form name
	newFileName := body.Name + filepath.Ext(uploadedImage.Filename)

	// save to folder
	if err := c.SaveUploadedFile(uploadedImage, "images/"+newFileName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal menyimpan gambar"})
		return
	}
	// save to database
	images := models.Images{
		Name: body.Name,
		Path: os.Getenv("URL_LINK") + "images/" + newFileName,
	}

	result := initializers.DB.Create(&images)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal memuat images"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": images})

}
