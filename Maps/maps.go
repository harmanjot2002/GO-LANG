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
	fmt.Println(a1) //map[]
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

	//Accessing Map Elements
	var a3 = make(map[string]string)
	a3["brand"] = "Ford"
	a3["model"] = "Mustang"
	a3["year"] = "1964"
	fmt.Printf(a3["brand"])

	//Updating and Adding Map Elements
	var a5 = make(map[string]string)
	a5["brand"] = "Ford"
	a5["model"] = "Mustang"
	a5["year"] = "1964"
	fmt.Println(a5)
	a5["year"] = "1970" // Updating an element
	a5["color"] = "red" // Adding an element
	fmt.Println(a5)

	//Remove Element from Map
	//Removing elements is done using the delete() function.
	var a7 = make(map[string]string)
	a7["brand"] = "Ford"
	a7["model"] = "Mustang"
	a7["year"] = "1964"
	fmt.Println(a7)
	delete(a7,"year")
	fmt.Println(a7)

	//Check For Specific Elements in a Map
	//If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.
	var a9 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}
	val1, ok1 := a9["brand"] // Checking for existing key and its value
	val2, ok2 := a9["color"] // Checking for non-existing key and its value
	val3, ok3 := a9["day"]   // Checking for existing key and its value
	_, ok4 := a9["model"]    // Only checking for existing key and not its value
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	/*
		Maps Are References
		Maps are references to hash tables.
		If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	*/
	var a89 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b89 := a89
	fmt.Println(a89)
	fmt.Println(b89)
	b89["year"] = "1970"
	fmt.Println("After change to b:")
	fmt.Println(a89)
	fmt.Println(b89)

	//Iterating Over Maps
	//You can use range to iterate over maps.
	a35 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range a35 {
		fmt.Printf("%v : %v, ", k, v)
	}

}