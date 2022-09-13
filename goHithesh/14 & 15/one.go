package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("hello")

	// basic declaration of slices
	// slices are more useful than arrays

	var data=[]string {"hello","there","this","is"}
	
	data=append(data,"knk","here" )
	fmt.Println(data)
	data=append(data[1:])
	fmt.Println(data)




	// another way of creating slices

	highscores:=make([]int,3)

	fmt.Println(highscores)
	highscores[0]=34
	highscores[1]=343
	highscores[2]=343324
	fmt.Println(highscores)
	// highscores[3]=34342
	// didn't work because the memory is not allocated

	highscores=append(highscores, 54,33,534)
	// append will automatically re allocate the memory
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)




// how to remove slice based on index


var courses=[] string {"c","c++","js","python"}
fmt.Println(courses)
// if i want to remove the js
var index int=2

courses=append(courses[:index],courses[index+1:]...)
// same as spread operator in js
// here ... allows us to combine these two arrays
fmt.Println(courses)














}