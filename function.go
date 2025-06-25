package main
import "fmt"

func rating(score float32) string {
	// score 0-5 return disappointing
	// score 6-7 return average
	// score 8-9 return good
	// score 10 return excellent 
	if(score<=6){
		return "disappointing"
	}
	if(score<=7){
		return "average"
	}
	if(score<=9){
		return "good"
	}
	if(score==10){
		return "excellent"
	}
	return "invalid score"
}

func main() {
	// This is a placeholder for the main function.
	// The actual implementation will be added later.
	//print 
 fmt.Println(rating(5)) // Output: disappointing
 fmt.Println(rating(6)) // Output: average
 fmt.Println(rating(8)) // Output: good
 fmt.Println(rating(10)) // Output: excellent
}