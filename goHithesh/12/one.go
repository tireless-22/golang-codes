package main

import "fmt"

func main() {

	fmt.Println("hello")
	var one=2
	fmt.Println(one)

	var ptr *int
	fmt.Println(ptr)
	// the output is nill
	// if the pointer is not initialized only declared then the value is nil

	myNumber:=34
	// it is very similar to c++
	var ptr1 *int
	ptr1=&myNumber
	// &=> reference
fmt.Println(*ptr1)
// same properties like we have in c++



}