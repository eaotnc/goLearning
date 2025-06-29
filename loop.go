package main

func main() {
	// create an array of movie genres
	var genres = [5]string{"Action", "Comedy", "Drama", "Horror", "Sci-Fi"}
	// print array length
	println("Number of genres:", len(genres))
	// print each genre using a for loop
	for i := 0; i < len(genres); i++ {
		println("Genre", i+1, ":", genres[i])
	}
	// print each genre using a range loop
	for index, genre := range genres {
		println("Genre", index+1, ":", genre)
	}
	// print each genre using a for loop with range
	for i := range genres {
		println("Genre", i+1, ":", genres[i])
	}
	// print each genre using a for loop with range and underscore
	for _, genre := range genres {	
		println("Genre:", genre)
	}
	// print each genre using a for loop with range and index
	for i, _ := range genres {
		println("Genre", i+1, ":", genres[i])
	}
	// print each genre using a for loop with range and index and underscore
	for i, genre := range genres {
		println("Genre", i+1, ":", genre)
	}
	// print each genre using a for loop with range and index and underscore
	for i := 0; i < len(genres); i++ {
		println("Genre", i+1, ":", genres[i])
	}
}
