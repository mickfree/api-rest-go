package handlers

import (
	"net/http"
	"strconv"

	"github.com/api-rest-go/database"
	"github.com/api-rest-go/models"

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
	title := c.PostForm("title")
	artist := c.PostForm("artist")
	year, _ := strconv.Atoi(c.PostForm("year"))
	genre := c.PostForm("genre")
	language := c.PostForm("language")
	duration, _ := strconv.ParseInt(c.PostForm("duration"), 10, 64)

	var coverImagePath string
	file, err := c.FormFile("cover")
	if err == nil {
		coverImagePath = "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, coverImagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}
	}

	album := models.Album{
		Title:      title,
		Artist:     artist,
		Year:       year,
		Genre:      genre,
		Language:   language,
		Duration:   duration,
		CoverImage: coverImagePath,
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

	// Solo actualiza si los campos están presentes
	if title := c.PostForm("title"); title != "" {
		album.Title = title
	}
	if artist := c.PostForm("artist"); artist != "" {
		album.Artist = artist
	}
	if yearStr := c.PostForm("year"); yearStr != "" {
		if year, err := strconv.Atoi(yearStr); err == nil {
			album.Year = year
		}
	}
	if genre := c.PostForm("genre"); genre != "" {
		album.Genre = genre
	}
	if language := c.PostForm("language"); language != "" {
		album.Language = language
	}
	if durationStr := c.PostForm("duration"); durationStr != "" {
		if duration, err := strconv.ParseInt(durationStr, 10, 64); err == nil {
			album.Duration = duration
		}
	}

	file, err := c.FormFile("cover")
	if err == nil {
		newPath := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, newPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar la imagen"})
			return
		}
		album.CoverImage = newPath
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

func SearchAlbums(c *gin.Context) {
	query := c.Query("q")
	var albums []models.Album

	if err := database.DB.Where("title ILIKE ?  or artist ILIKE ? or CAST(year AS TEXT) ILIKE ? or genre ILIKE ? or language ILIKE ?", "%"+query+"%","%"+query+"%","%"+query+"%","%"+query+"%","%"+query+"%").Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error al buscar albumes"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func FilterAlbums(c *gin.Context) {
	title := c.Query("title")


	db := database.DB


	if title != "" {
		db = db.Where("title ILIKE ?", "%"+title+"%")
	}

	var albums []models.Album
	if err := db.Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"No encontro el elemento Indicado"})
		return
	}
	c.JSON(http.StatusOK, albums)
}