package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Artist   string  `json:"artist,omitempty"`
	Price    float64 `json:"price"`
	Password string  `json:"-"`
}

var albums = []album{
	{ID: "1", Title: "Blue", Artist: "Jho asan", Price: 16.66},
	{ID: "2", Title: "Green", Artist: "Mariah", Price: 25.99},
	{ID: "3", Title: "Yellow", Artist: "Ken", Price: 83.22},
}

func main() {
	serviceName := os.Getenv("service")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, serviceName)
	})

	router.Run()
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})

}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
