package main

import "fmt"

func main() {
	// arrays is not much used thing in golang
// slices is most used thing in the golang


	fmt.Println("hello there")

	var fruits[4] string

	fruits[0]="apple"
	fruits[1]="banana"
	fruits[3]="papaya"
	fmt.Println("fruits are",fruits)
	// [apple banana  papaya]
	// we did not put anything in the fruits[2] so that is not printed as any garbage value 
		fmt.Println("fruits are",len(fruits))




		var vegList=[3] string{"potato","beans","mushroom"}
		fmt.Println(vegList)
}