package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")

	welcome :="welcome to the user input"
	fmt.Println(welcome)

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for out Pizza")

	// comma ok || err ok syntax

	input,err :=reader.ReadString('\n')
	fmt.Println("entered values is ",input)

	fmt.Println(err)




}