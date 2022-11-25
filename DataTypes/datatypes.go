package main

import (
	"fmt"
)

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String
  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)

  //Signed Integers
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Print("\n")
  fmt.Printf("Type: %T, value: %v", y, y)

  //Unsigned Integers
  var x1 uint = 500
  var y1 uint = 4500
  fmt.Printf("Type: %T, value: %v", x1, x1)
  fmt.Printf("Type: %T, value: %v", y1, y1)

  //Float32
  var x3 float32 = 123.78
  var y3 float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x3, x3)
  fmt.Printf("Type: %T, value: %v", y3, y3)

  //Float64
  var x4 float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", x4, x4)
}