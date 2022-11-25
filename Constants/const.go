package main

import ("fmt")

//Enumerated Constants
const (
	x=iota+2
	y=iota
	z=iota
)

const (
	a=iota
	b=iota
)

const (
	x1=iota
	_
	y1
	z1
)

func main() { 
	//The value of a constant must be assigned when you declare it.
	const PI = 3.14
	fmt.Println(PI) //3.14

	//  Typed Constants
	// Typed constants are declared with a defined type
	const A int=1
	fmt.Println(A) //1
	// Untyped Constants
	// Untyped constants are declared without a type

	//Enumerated constants make use of iota and it automatically understands that we wish to increment its value by 1.
	//It increments value of iota ,not x when we are calling y and z here
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	//Value of iota applies to its own group only,in next group it will again start from 0.
	fmt.Println(a)
	fmt.Println(b)
	//Here underscore(_) takes place of 1.
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)
}