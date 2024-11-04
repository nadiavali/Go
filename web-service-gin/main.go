package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"errors"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func createAlbums(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err!= nil{
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getalbumbyid(c *gin.Context){
	id := c.Param("id")
	album, err := getAlbumsById(id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, album)

}

func getAlbumsById(id string)(*album, error){
	for i, a := range albums{
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("album not found")
}

func main() {
    router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
    router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbums)
	router.GET("/albums/:id", getalbumbyid)
    router.Run("localhost:8080")
}

