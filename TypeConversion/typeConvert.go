package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int32 = 100
	var b int64 = int64(a)
	//Writing only "var b int64=a" donot work
	fmt.Printf("%v,%T", b, b)
	fmt.Printf("\n")

	//To convert one data type to some other different data type,we need to include package "strconv" and use "Itoa" function of it.
	var str  int=65
	var d string=strconv.Itoa(str)
	fmt.Printf("%v,%T",d,d)
	fmt.Printf("\n")

	//How to print ASCII values.
	var e string=string(str)
	fmt.Printf("%v,%T",e,e)
	fmt.Printf("\n")

	//If we donot specify that we want to take 32 bit or 64 bit float,then it will go with 64 bit by default,i.e,it goes with maximum sized one.
	c := 2.5
	fmt.Printf("%v,%T", c, c)
	fmt.Printf("\n")
}
