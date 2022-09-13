package main

import (
	"fmt"

	
)

func main() {

	// maps in golang



	fmt.Println("hello")
	// make is more useful than new in golang

	languages:=make(map[string]string)

	fmt.Println(languages)

	languages["JS"]="JavaScript"
	languages["RB"]="Ruby"
	languages["PY"]="Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

// LOOPS

for key,value:=range languages{
	fmt.Println("key",key,"value",value)
}

// valarus operator will take care of the comma ok syntax











}