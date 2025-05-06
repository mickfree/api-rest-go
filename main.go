package main

import (
	"github.com/api-rest-go/database"
	"github.com/api-rest-go/handlers"
	"github.com/api-rest-go/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	database.Connect()

	// migrar automaticamente
	database.DB.AutoMigrate(&models.Album{})
	// Endpoints
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbum)
	router.POST("/albums", handlers.CreateAlbum)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)
	router.Static("/uploads", "./uploads")
	router.GET("/albums/search", handlers.SearchAlbums)
	router.GET("/albums/filter", handlers.FilterAlbums) /// albums/filter?artist=floyd&genre=rock

	// Port
	router.Run(":8888")

}
