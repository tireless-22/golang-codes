package main

import "fmt"



		// numbeOfUsers:=300000.0
		// /you can not use this syntax here


const loginToken string ="helloToken"



func main() {

	var username string = "naveen karthik"

	fmt.Print(username)
	fmt.Printf("variable is of type : %T \n",username)


	var isLoggedIn bool =true
		fmt.Printf("variable is of type : %T \n",isLoggedIn)

		var intValue int 
		fmt.Printf("variable is of type: %T",intValue)
		fmt.Print(intValue)


			var stringValue string
		fmt.Printf("variable is of type: %T",stringValue)
		fmt.Print(intValue)
		// implicit type

		var website="learncodeonline.in"
		fmt.Print(website)
		// website=3
		// is not possible

		// no var style
// In Go, := is for declaration + assignment, whereas = is for assignment only.
		numbeOfUsers:=300000.0
		fmt.Print(numbeOfUsers)

		fmt.Println(loginToken)







	}