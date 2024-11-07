package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct{
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Quantity float64 `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Be" , Author: "Nadia Vali", Quantity: 3},
	{ID: "2", Title: "Be Smart" , Author: "Nadia Vali", Quantity: 5},
	{ID: "3", Title: "Be Patient" , Author: "Nadia Mirali", Quantity: 9},
	{ID: "4", Title: "Be Sure" , Author: "Nadia Vali", Quantity: 4},
} 

func GetAllBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func Createbook(c *gin.Context){
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookById(c *gin.Context){
	id := c.Param("id")
	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
}

func DeleteBook(c *gin.Context){
	id := c.Param("id")
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"messegae": fmt.Sprintf("Book %s by %s has been deleted", b.Title, b.Author)})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func UpdateBookUpdate(c *gin.Context){
	id := c.Param("id")
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid json data"})
		return
	}

	for i, b := range books{
		if b.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book updated", "book": books[i]})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func UpdatebyTitle(c *gin.Context){
	id := c.Param("id")
	var bookUpdated map[string]interface{}
	if err := c.BindJSON(&bookUpdated); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid json data"})
		return
	}

	for i, b := range books{
		if b.ID == id{
			if title, ok := bookUpdated["title"].(string); ok{
				books[i].Title = title
			}
			if author, ok := bookUpdated["author"].(string); ok {
				books[i].Author = author
			}
			c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": books[i]})
            return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

func main(){
	router := gin.Default()
	router.GET("/books", GetAllBooks)
	router.POST("/books", Createbook)
	router.GET("/books/:id", GetBookById)
	router.DELETE("/books/:id", DeleteBook)
	router.PUT("/books/:id", UpdateBookUpdate)
	router.PATCH("/books/:id", UpdatebyTitle)
	router.Run("localhost:8080")
}




//package main

// import (
// 	"net/http"
// 	"strconv"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// )

// Movie represents each movie in our list
type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}

// In-memory slice to store movies and a mutex for safe concurrent access
var (
	movies []Movie
	mu     sync.Mutex
	nextID = 1
)

// POST /movies: Add a new movie
func createMovie(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	movie.ID = nextID
	nextID++
	movies = append(movies, movie)
	c.JSON(http.StatusCreated, movie)
}

// GET /movies: Retrieve all movies
func getMovies(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, movies)
}

// GET /movies/:id: Retrieve a specific movie by ID
func getMovieByID(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	for _, movie := range movies {
		if movie.ID == id {
			c.JSON(http.StatusOK, movie)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

// PUT /movies/:id: Update a movie by ID
func updateMovie(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			if err := c.ShouldBindJSON(&movies[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
				return
			}
			movies[i].ID = id // Preserve the original ID
			c.JSON(http.StatusOK, movies[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

// DELETE /movies/:id: Delete a movie by ID
func deleteMovie(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{"message": "Movie deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

//func main() {
// 	r := gin.Default()

// 	// Define routes
// 	r.POST("/movies", createMovie)
// 	r.GET("/movies", getMovies)
// 	r.GET("/movies/:id", getMovieByID)
// 	r.PUT("/movies/:id", updateMovie)
// 	r.DELETE("/movies/:id", deleteMovie)

// 	r.Run(":8080") // Start the server on port 8080
// }
