package main

import (
	"fmt"
	"time"


)

func main() {

	fmt.Println("hello welcome to time study")
	presentTime:=time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))



	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:=time.Date(2002,time.February,22,22,22,0,0,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006"))

}