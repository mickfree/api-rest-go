package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Artist   string        `json:"artist"`
	Year     int           `json:"year"`
	Genre    string        `json:"genre"`
	Language string        `json:"language"`
	Duration time.Duration `json:"duration"`
}

var albums = []album{
    {
        ID:       "1",
        Title:    "Thriller",
        Artist:   "Michael Jackson",
        Year:     1982,
        Genre:    "Pop",
        Language: "English",
        Duration: 42*time.Minute + 19*time.Second,
	},
    {
        ID:       "2",
        Title:    "The Dark Side of the Moon",
        Artist:   "Pink Floyd",
        Year:     1973,
        Genre:    "Progressive Rock",
        Language: "English",
        Duration: 42*time.Minute + 49*time.Second,
    },
    {
        ID:       "3",
        Title:    "Bohemian Rhapsody",
        Artist:   "Queen",
        Year:     1975,
        Genre:    "Rock",
        Language: "English",
        Duration: 5*time.Minute + 55*time.Second,
    },
}

func getAlbum(c *gin.Context){
	id := c.Param("id")

	for _, a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})

}

// no necesita control de error por que estoy llamando una var global
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(c *gin.Context){
	var createAlbum album

	if err := c.BindJSON(&createAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Json no creado"})
		return
	}

	albums = append(albums, createAlbum)
	c.IndentedJSON(http.StatusCreated, albums)

}


func updateAlbum(c *gin.Context) {
    id := c.Param("id")
    var updatedAlbum album

    if err := c.BindJSON(&updatedAlbum); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "JSON inválido"})
        return
    }

    for i, a := range albums {
        if a.ID == id {
			updatedAlbum.ID = id
            albums[i] = updatedAlbum
            c.IndentedJSON(http.StatusOK, updatedAlbum)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message":"Album Eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album no encontrado"})
}

func main() {
	// Endpoints
	router := gin.Default()
	router.GET("/albums/:id", getAlbum)
	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)
	// Port
	router.Run("localhost:8888")
}
