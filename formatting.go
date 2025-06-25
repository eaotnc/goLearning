package main

import "fmt"
func main() {
	msg, age, price, check := "Hello Gopher!!!", 55, 22.52,	true
	var r rune = '😍'	
	fmt. Printf("type: %T -- msg: %#v\n",msg, msg)
	fmt. Printf("type: %T -- age: %#v\n",age, age)
	fmt. Printf("type: %T -- price: %#v\n", price, price) 
	fmt. Printf("type: %T -- check: %#v\n", check, check)
	fmt. Printf("type: %T - r: %#v\n", r, r)

}