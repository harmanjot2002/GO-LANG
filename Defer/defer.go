package main

import ("fmt")

func main(){
	//The defer keyword is used to delay the execution of a function or a statement until the nearby function returns. In simple words, defer will move the execution of the statement to the very end inside a function.
	defer fmt.Println("I")
	fmt.Println("love")
	fmt.Println("Programming")

	fmt.Println("I")
	defer fmt.Println("love")
	fmt.Println("Programming")

	fmt.Println("I")
	fmt.Println("love")
	defer fmt.Println("Programming")

	defer fmt.Println("I")
	defer fmt.Println("love")
	defer fmt.Println("Programming")
}