package main

import "fmt"

func main() {

	


	// structs are alternative to classes in golang
	// no inheritance , no super or parent

	naveen:=User{"naveen","knk@gmail.com",false,23}
	fmt.Println(naveen)
	fmt.Printf("the details are %+v\n",naveen)
	fmt.Print(naveen.Name)
	fmt.Printf(naveen.Email)

}

// first capital indicated public

type User struct{
	Name string
	Email string
	Status bool
	Age int

}




