package main

import ("net/http"
		"github.com/gin-gonic/gin"
		"fmt"
)

type Movie struct{
	ID	string	`json:"id"`
	Film string `json:"film"`
}


var movies = []Movie{
{ID:"1", Film: "Cloud Atlas"},
{ID:"2", Film: "Shutter Island"},
}


func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}


func getMovieByID(c *gin.Context) {
	id := c.Param("id")
	for _, m := range movies {
		if m.ID == id {
			c.IndentedJSON(http.StatusOK, m)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie Not Found"})
}

func CreateMovies(c *gin.Context) {
	var newMovie Movie
	if err := c.BindJSON(&newMovie); err != nil{
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)

}


func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	for i, m := range movies {
		if m.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Movie %s deleted", m.Film)})
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Movie Not Found"})
}



func updatedMovie(c *gin.Context){
	id := c.Param("id")
	var updatedMovie Movie

	if err := c.BindJSON(&updatedMovie); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "some"})
		return
	}

	for i, m := range movies {
		if m.ID == id {
			updatedMovie.ID = m.ID
			movies[i] = updatedMovie
			c.IndentedJSON(http.StatusOK, gin.H{"message":"updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error":"some"})
}



func main() {
	router := gin.Default()
	router.GET("/movies", GetMovies)
	router.GET("/movies/:id", getMovieByID)
	router.POST("/movies", CreateMovies)
}