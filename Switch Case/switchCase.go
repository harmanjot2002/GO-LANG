package main

import ("fmt")

func main(){
	//The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.
	switch 5{
	case 5:
		fmt.Println("Hi")
	case 4:
		fmt.Println("Hello")
	default:
		fmt.Println("None")
	}

	//Using fallthrough also executes next statement in additional to desired one.
	switch 5{
	case 5:
		fmt.Println("Hi")
		fallthrough
	case 4:
		fmt.Println("Hello")
	default:
		fmt.Println("None")
	}

	//Go supports multi-case switch statements.
	day := 5
   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}