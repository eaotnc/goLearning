package main

import "fmt"

func main(){
     
	// create slice of user lating for move in 2 slices 
	var userRatings = []int{5, 4, 3, 2, 1}
	var userRatings2 = []int{1, 2, 3, 4, 5}

	// combine the two slices into one
	var combinedRatings = append(userRatings, userRatings2...)
	// print all combined ratings
	// using fmt.Printf for formatted output
	fmt.Printf("combinedRatings: %v\n", combinedRatings)
	fmt.Printf("combinedRatings: %#v\n", combinedRatings)
	// show combined ratings 5-8 using slice syntax
	fmt.Printf("combinedRatings[5:8]: %v\n", combinedRatings[5:8])


	// create empty slice of strings
	var emptySlice = make([]string,3)
	emptySlice[1] = "Hello" // This will cause a runtime error because the slice is empty and has no capacity
	fmt.Printf("emptySlice: %v\n", emptySlice)
}
