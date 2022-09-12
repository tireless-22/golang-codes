package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("hello there 7")

	reader:=bufio.NewReader(os.Stdin)

	// comma ok || err err

	input,_:=reader.ReadString('\n')

	fmt.Println(input)
	fmt.Printf("the rating is %T",input)


	

}