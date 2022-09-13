package main

import "fmt"

func main() {

	fmt.Println("hello")

	vals:=[] string {"hello","there"}

	fmt.Println(vals)

	for ind:=0;ind<len(vals);ind++{
		fmt.Println(vals[ind])

	}
	fmt.Println("")


	for i := range(vals){
		// if we try to put only one this will give the indexes only
		fmt.Println(vals[i])
	}

	fmt.Println("")


	for i,j:=range(vals){
		// if we try to get 2 values then we will get the index and value
		fmt.Println(i,j)
	}

	fmt.Println("")

	for _,i:=range(vals){
		// if we did n't want to use the indexes then put _ there 
		// if we didn't put the _ then it will give us error
		fmt.Println(i)

	}

	// 


	// it is just like a while loop in golang

	ind:=1
	for ind<10{
		fmt.Println(ind)
		ind++
	}


	// only i++
	// ++i is not supported in golang

	 







}