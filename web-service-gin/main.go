package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// getAlbums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums,)
}

//getAlbumByID 
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//loop over the list of albums, looking for the required ID
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
}

var albums = []album{
	{ID: "1", Title: "BLue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id",getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
