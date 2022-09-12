package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello 8")
	reader:=bufio.NewReader(os.Stdin)

	// we have to handle the error here only by using comma ok syntax

	val,_:=reader.ReadString('\n')

	fmt.Printf("the rating is of type %T  ",val)

	fmt.Println("performing string convertion")


	// numRating ,err:=strconv.ParseFloat(val,64)
	// if we did not remove the spaces after the value we will get the below error
// 	the rating is of type string  performing string convertion
// strconv.ParseFloat: parsing "4\r\n": invalid syntax

numRating,err:= strconv.ParseFloat(strings.TrimSpace(val),64);
// we have to use the TrimSpace function of strings library in golang


// we have to check if there is any error there
// if there is no error the default value is nil
	if err!= nil{
		fmt.Println(err)

	}else{
		fmt.Println("check")
		fmt.Println(numRating)
	}





	fmt.Println(val)




}