package main

import ("fmt")

//Global Level Variables
var val2 int = 100
//Package Level Variables
var myValue int = 500

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
  //:= this notation of declaring keywords cannot be used to declare variables in global scope,for that purpose we can only use "var" keyword.
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
 //We can declare variables in three scopes:
 //1.Local Level-inside a function.
 //2.Global Level-outside function,but condition is to always write it in Pascal Case. 
 //3.Package Level-outside a function,but condition is to always write it in Camel Case.
 //Package Level=Global+Will be accessed in all other packages also
  fmt.Println(val2) //100
  fmt.Println(myValue) //500
  //Shadowing
  //It means on initializing same variable name at global and local level,priority will be given to local variable only.
  val2=55
  fmt.Println(val2) //55


 //Constants
 //The value of a constant must be assigned when you declare it.
 const PI = 3.14
 fmt.Println(PI) //3.14
//  Typed Constants
// Typed constants are declared with a defined type:
const A int=1
fmt.Println(A) //1
// Untyped Constants
// Untyped constants are declared without a type



}