package main

import ("fmt")

func main(){
	//Go only uses "For Loop"
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}

	for i:=0; i <= 100; i+=10 {
		fmt.Println(i)
	}

	for i:=0; i < 5; i++ {
		if i == 3 {
		  continue
		}
	   fmt.Println(i)
	}

	for i:=0; i < 5; i++ {
		if i == 3 {
		  break
		}
	   fmt.Println(i)
	}

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i < len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
		fmt.Println(adj[i],fruits[j])
		}
	}

	fruits1 := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits1 {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	fruits2 := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits2 {
		fmt.Printf("%v\n", val)
	}
}