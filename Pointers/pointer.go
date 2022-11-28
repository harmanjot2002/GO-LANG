package main

import ("fmt")

func main(){
	var x int=90
	var b *int=&x
	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(*b) //Dereferencing
}