package main

import ("fmt")

func main(){
	/*
	Maps are:
		1.key-value pairs
		2.unchangable
		3.unordered
		4.default value is "nil"
		5.it holds references to underlying hash table
		6. The order of the map elements defined in the code is different from the way that they are stored. The data are stored in a way to have efficient data retrieval from the map.
	*/

	//Creating Maps Using var and :=
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	
	//Creating Maps Using Using make()Function:
	var a1 = make(map[string]string) // The map is empty now
	a1["brand"] = "Ford"
	a1["model"] = "Mustang"
	a1["year"] = "1964"
	// a is no longer empty
	b1 := make(map[string]int)
	b1["Oslo"] = 1
	b1["Bergen"] = 2
	b1["Trondheim"] = 3
	b1["Stavanger"] = 4
	fmt.Printf("a1\t%v\n", a1)
	fmt.Printf("b1\t%v\n", b1)
	
	/*
		Creating an Empty Map
		There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.

		The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.
	*/
	var a2 = make(map[string]string)
	var b2 map[string]string
	fmt.Println(a2 == nil)
	fmt.Println(b2 == nil)
}