/*Filter and Sort Data from a List
Challenge: Write a function that takes a list of movies (with fields like Title, Genre, and Year)
and returns a filtered list based on genre and sorted by release year.
Skills Tested: Filtering and sorting, slice manipulation.
Hints: Use Go’s built-in sort package for sorting.*/

package main

import (
	"fmt"
	"sort"
)

// Define the Movie struct
type Movie struct {
	Title string
	Genre string
	Year  int
}

// filterAndSortMovies filters movies by genre and sorts the results by year
func filterAndSortMovies(movies []Movie, genre string) []Movie {
	// Step 1: Filter movies by the specified genre
	var filteredMovies []Movie
	for _, movie := range movies {
		if movie.Genre == genre {
			filteredMovies = append(filteredMovies, movie)
		}
	}

	// Step 2: Sort the filtered movies by year (ascending)
	sort.Slice(filteredMovies, func(i, j int) bool {
		return filteredMovies[i].Year < filteredMovies[j].Year
	})

	return filteredMovies
}

func main() {
	// Sample list of movies
	movies := []Movie{
		{Title: "Inception", Genre: "Sci-Fi", Year: 2010},
		{Title: "The Matrix", Genre: "Sci-Fi", Year: 1999},
		{Title: "Interstellar", Genre: "Sci-Fi", Year: 2014},
		{Title: "The Godfather", Genre: "Crime", Year: 1972},
		{Title: "Pulp Fiction", Genre: "Crime", Year: 1994},
		{Title: "The Dark Knight", Genre: "Action", Year: 2008},
	}

	// Filter by genre "Sci-Fi" and sort by year
	filteredAndSortedMovies := filterAndSortMovies(movies, "Sci-Fi")

	// Print the results
	fmt.Println("Filtered and Sorted Movies:")
	for _, movie := range filteredAndSortedMovies {
		fmt.Printf("Title: %s, Genre: %s, Year: %d\n", movie.Title, movie.Genre, movie.Year)
	}
}
