package main

import ("fmt")
 
func main(){
	//Fixed length Arrays
	var arr [4]int = [4]int{1,2,3,4}
	var arr1 = [3]int{1,2,3}
  	arr2 := [5]int{4,5,6,7,8}
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	//Inferred length Arrays
	var arr11 = [...]int{1,2,3}
	arr21 := [...]int{4,5,6,7,8}
	fmt.Println(arr11)
	fmt.Println(arr21)

	//Arrays of strings
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  	fmt.Print(cars)

	//Access elements of array
	prices := [3]int{10,20,30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//Change Elements of an Array
	prices2 := [3]int{10,20,30}
	prices2[2] = 50
	fmt.Println(prices2)

	/*
		Array Initialization
		If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
		Tip: The default value for int is 0, and the default value for string is "".
	*/
	arr12 := [5]int{} //not initialized
	arr22 := [5]int{1,2} //partially initialized
	arr32 := [5]int{1,2,3,4,5} //fully initialized
	fmt.Println(arr12)
	fmt.Println(arr22)
	fmt.Println(arr32)	

	/*
	Initialize Only Specific Elements
	It is possible to initialize only specific elements in an array.
	*/
	arr13 := [5]int{1:10,2:40}
  	fmt.Println(arr13)

	/*
	Find the Length of an Array
	The len() function is used to find the length of an array
	*/
	arr14 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr24 := [...]int{1,2,3,4,5,6}
	fmt.Println(len(arr14))
	fmt.Println(len(arr24))

	//Copy elements of one array into another
	var arr15 =[...]int{1,2,3,4,5,6,7,8,9}
	arr15[1]=41
	arr16 := arr15
	arr16[2]=51
	fmt.Println(arr15)
	fmt.Println(arr16)

	//multi-dimensional arrays
	var matrix [3][3]int = [3][3] int {{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(matrix)
	fmt.Println(len(matrix))
	fmt.Println(matrix[1])
	
}