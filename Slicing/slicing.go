package main

import ("fmt")

func main(){
	//unlike arrays, the length of a slice can grow and shrink as you see fit.
	/*
		In Go, there are several ways to create a slice:
		Using the []datatype{values} format
		Create a slice from an array
		Using the make() function
	*/

	//Declare an empty slice of 0 length and 0 capacity.
	myslice := []int{}
	fmt.Println(myslice)

	myslice2 := []int{1,2,3}
	fmt.Println(myslice2)

	/*
		In Go, there are two functions that can be used to return the length and capacity of a slice:
		len() function - returns the length of the slice (the number of elements in the slice)
		cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	*/
	myslice12 := []int{}
	fmt.Println(len(myslice12))
	fmt.Println(cap(myslice12))
	fmt.Println(myslice12)

	myslice21 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice21))
	fmt.Println(cap(myslice21))
	fmt.Println(myslice21)

	//Create a slice from an array
	arr13 := [6]int{10, 11, 12, 13, 14,15}
  	myslice3 := arr13[2:4]
	fmt.Printf("myslice = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
	//If myslice3 started from element 0, the slice capacity would be 6.

	//Create a Slice With The make() Function
	//Syntax:
	//slice_name := make([]type, length, capacity)
	//If the capacity parameter is not defined, it will be equal to length.
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	// with omitted capacity
	myslice25 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice25)
	fmt.Printf("length = %d\n", len(myslice25))
	fmt.Printf("capacity = %d\n", cap(myslice25))

	//Access Elements of a Slice
	prices34 := []int{10,20,30}
	fmt.Println(prices34[0])
	fmt.Println(prices34[2])
}