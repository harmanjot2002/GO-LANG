package main

import ("fmt")

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func familyName1(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction(x int, y int) int {
	return x + y
}
//int after parenthesis of parameters specify return type of function

func myFunction1(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction2(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func summation(vals ...int){
	sum:=0
	for _,n := range vals{
		sum+=n
	}
	fmt.Println(sum)
}

func main() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
	//When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.

	familyName1("Liam", 3)
	familyName1("Jenny", 14)
	familyName1("Anja", 30)
	//When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.

	fmt.Println(myFunction(1, 2))

	total := myFunction1(1, 2)
  	fmt.Println(total)
	//You can also store the return value in a variable

	_, b := myFunction2(5, "Hello")
  	fmt.Println(b)
	//If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.

	summation(2,3,4,5,6)
	//When we are not sure that how many parameters we wish to pass,we use ...
}
