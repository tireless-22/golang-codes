package main

import "fmt"

// jwtToken:="hello"
// valarus operator is not valid here

var jwtToken="hello"
// other two types are allowed here

const loginToken string="sdfasd";


func main() {

	var username string="knk"
	// pf shortcut for fmt.println("")
	fmt.Println(username)

	var isLoggedIn bool=false
	fmt.Printf("variable is of type: %T",isLoggedIn)


	fmt.Println(isLoggedIn)



	// default values and some variables

	var anotherOne int
	fmt.Println(anotherOne)
	// 0 is the defualt of integer

		var anotherStr string
	fmt.Println(anotherStr)
	// 0 is the defualt of integer

	//========> another style of declaring variable, we need not to put the varibale type
	// it will decided if we did not put the type
	// but we are not allowed to change that


	var someName="knk"

	fmt.Println(someName)

	// someName=2
	// it will throw error


	// =========> no var style
	// it is called valarus operator
	// we can this inside any motheds 
	// we are not allowed to use this outside of methods or in the global scope


	numberOfUser:=343453
	fmt.Println(numberOfUser)


	fmt.Println(loginToken)
















}