/*2. Fetch and Process External JSON Data
Challenge: Simulate fetching JSON data from a fake API (you can just define the data in a variable for simplicity) and convert it into a Go struct, then display it in a formatted response.
Skills Tested: JSON parsing and struct usage.
Hints: Keep the JSON data simple and focus on encoding/decoding.*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define the structure of the data we're simulating
type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}

// Simulated JSON data as if it were fetched from an external API
var jsonData = `
[
	{"id": 1, "title": "Inception", "genre": "Sci-Fi", "year": 2010},
	{"id": 2, "title": "The Matrix", "genre": "Sci-Fi", "year": 1999},
	{"id": 3, "title": "Interstellar", "genre": "Sci-Fi", "year": 2014}
]`

// fetchMovies simulates fetching JSON data and decoding it into a Go struct
func fetchMovies() ([]Movie, error) {
	var movies []Movie

	// Decode JSON data into the movies slice
	err := json.Unmarshal([]byte(jsonData), &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// moviesHandler handles HTTP requests and displays the fetched movies in JSON format
func moviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := fetchMovies()
	if err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		log.Println("Error fetching movies:", err)
		return
	}

	// Set response header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	// Encode the movies slice into JSON and write it to the response
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Setup HTTP server with /movies endpoint
	http.HandleFunc("/movies", moviesHandler)

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
