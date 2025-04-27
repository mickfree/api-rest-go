package handlers

import (
	"github.com/api-rest-go/database"
	"github.com/api-rest-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// no necesita control de error por que estoy llamando una var global
func GetAlbums(c *gin.Context) {
	var albums []models.Album
	database.DB.Find(&albums)
	c.JSON(http.StatusOK, albums)
}

// cuando provienen de un packete la funcion empieza con Mayuscula
func GetAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := database.DB.First(&album, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func CreateAlbum(c *gin.Context) {
	var album models.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON invalido"})
		return
	}

	database.DB.Create(&album)
	c.JSON(http.StatusCreated, album)

}

func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album

	if err := database.DB.First(&album, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON invalido"})
		return
	}

	database.DB.Save(&album)
	c.JSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	var album models.Album
	if err := database.DB.First(&album, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
		return
	}
	if err := database.DB.Delete(&models.Album{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al eliminar el album"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album eliminado"})
}
