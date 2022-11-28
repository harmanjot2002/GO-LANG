package main

import ("fmt")

type Employee struct{
	EmpId string
}

func main(){
	var x int=90
	var b *int=&x
	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(*b) //Dereferencing

	var emp *Employee
	fmt.Println(emp)
	emp=new(Employee)
	fmt.Println(emp)
}