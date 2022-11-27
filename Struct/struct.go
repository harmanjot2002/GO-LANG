package main

import ("fmt")

type Employee struct{
	EmpId int
	EmpName string
	EmpMobile []string
}

func main(){
	emp := Employee{
		EmpId:101,
		EmpName:"Ravi",
		EmpMobile:[] string{"1234567890","9876543210"},
	}
	emp2 := emp
	emp2.EmpName="Rahul"
	fmt.Println(emp)
	fmt.Println(emp2)
	// Both emp1 and emp2 has diff. addresses.
}