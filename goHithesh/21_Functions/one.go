package main

import "fmt"

func main() {
	fmt.Println("hello")
	greeter();

}

// we are not allowed to write a function inside another function


func greeter(){
	fmt.Println("namasthe")
}


func adder(valOne int, valTwo int)int {
	return valOne+valTwo
}

func proAdder(values ...int) int {
	// values is of type slice
	

	total:=0
	for _,val:=range values{
		total+=val
	}

	return total
}







