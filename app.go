package main

import ("fmt")

func main() { 
  //Printing Hello World
  fmt.Println("Hello World!")

  //Variables:
  //Declaration of variable-If we have declared some variable,then its necessary to use it too.
  var a int
  fmt.Println(a)
  //Initialization of Variable
  a=55
  fmt.Println(a)
  //Declaration with Initialization
  var b int=500
  fmt.Println(b)
  //More verbose method 
  var c=100
  fmt.Println(c)
  //Type Inference
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred
  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
  //Types of variables types
  var a1 string //" "
  var b1 int //0
  var c1 bool //false
  fmt.Println(a1)
  fmt.Println(b1)
  fmt.Println(c1)
  //Variable declaration in a Block
  var (
    a2 int
    b2 int = 1
    c2 string = "hello"
  )

 fmt.Println(a2)
 fmt.Println(b2)
 fmt.Println(c2)
  
}