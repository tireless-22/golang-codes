package main

import (
	"fmt"
	"time"

	
)

func main() {

	fmt.Println("hello")

	presentTime:= time.Now();
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"));
	// We have to put that time only
	

}
