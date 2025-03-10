package main 

import "fmt"

var msg string = "Hello, World!Eao"
func main() {
    msgWithoutVar := "Hello, msgWithoutVar"
	var age int = 20
	var name string = "John"
	var isCool bool = true
	var size float32 = 2.3

	age2,name2,isCool2 := 30,"John2",false

	fmt.Println("msg:",msg)
	fmt.Println("msgWithoutVar:",msgWithoutVar)
	fmt.Println("age:",age)
	fmt.Println("name:",name)
	fmt.Println("isCool:",isCool)
	fmt.Println("size:",size)

	fmt.Println("age2:",age2)
	fmt.Println("name2:",name2)
	fmt.Println("isCool2:",isCool2)

	fmt.Printf("%T\n:", msg)
	fmt.Printf("%T\n:", age)
	fmt.Printf("%T\n:", name)
	fmt.Printf("%T\n:", isCool)
	fmt.Printf("%T\n:", size)

	var typeOfSize = fmt.Sprintf("%T", size)
	fmt.Println("typeOfSize",typeOfSize)

}