package main

import ("fmt")

func main() {
	//%v:Value
	//%T:Type 
	//%%:Percentage symbol
	//%#v:Prints value in Go-syntax format
	var val=100
	fmt.Printf("%v, %T", val, val) //100,int
	var i = 15.5
  	var txt = "Hello World!"
	fmt.Print("\n")
	fmt.Printf("%v\n", i) //15.5
	fmt.Printf("%#v\n", i) //15.5
	fmt.Printf("%v%%\n", i) //15.5%
	fmt.Printf("%T\n", i) //float64
	fmt.Printf("%v\n", txt) //Hello World!
	fmt.Printf("%#v\n", txt) //"Hello World!"
	fmt.Printf("%T\n", txt) //string
}
