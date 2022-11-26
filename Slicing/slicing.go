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

	//Change Elements of a Slice
	prices43 := []int{10,20,30}
	prices43[2] = 50
	fmt.Println(prices43[0])
	fmt.Println(prices43[2])

	//Append Elements To a Slice
	//You can append elements to the end of a slice using the append()function:
	//slice_name = append(slice_name, element1, element2, ...)
	myslice17 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice17 = %v\n", myslice17)
	fmt.Printf("length = %d\n", len(myslice17))
	fmt.Printf("capacity = %d\n", cap(myslice17))
	//After appending:
	myslice17 = append(myslice17, 20, 21)
	fmt.Printf("myslice17 = %v\n", myslice17)
	fmt.Printf("length = %d\n", len(myslice17))
	fmt.Printf("capacity = %d\n", cap(myslice17))

	//Append One Slice To Another Slice
	//Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	myslice18 := []int{1,2,3}
	myslice28 := []int{4,5,6}
	myslice38 := append(myslice18, myslice28...)
	fmt.Printf("myslice38=%v\n", myslice38)
	fmt.Printf("length=%d\n", len(myslice38))
	fmt.Printf("capacity=%d\n", cap(myslice38))

	//Change The Length of a Slice
	//Unlike arrays, it is possible to change the length of a slice.
	//By reslicing or appending items 

	/*
		Memory Efficiency
		When using slices, Go loads all the underlying elements into the memory.

		If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

		The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 

		The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	*/
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers) 
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}