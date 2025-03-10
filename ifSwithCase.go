package main

import ("fmt";"math")

func main () {

	num := 2	
	if num%2 == 0 {
		fmt.Println("เลขคู่" )
	} else if num == 35 {
		fmt.Println("nufalagusan" )
	} else {
		fmt.Println("เลขโสด")
	}

	limit := 225.0
	if v := math.Pow(10,2); v < limit {
	fmt.Println("x power n is:", v)
	}else{
	fmt.Println("x power n is:", v)
	}
}