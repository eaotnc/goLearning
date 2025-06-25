package main

import "fmt"

func main() {
	// create an array of movie genres
	var genres = [5]string{"Action", "Comedy", "Drama", "Horror", "Sci-Fi"}

	// print array length
	fmt.Printf("Number of genres: %d\n", len(genres))
	// print each genre using a for loop
	for i := 0; i < len(genres); i++ {
		fmt.Printf("Genre %d: %s\n", i+1, genres[i])
	}
	

	// 
}