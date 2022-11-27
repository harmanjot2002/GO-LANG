package main

import ("fmt")

func main(){
	//In Go language,we have to start else block in same line where we ended } of if statement.
	//If-else if -else statement are used in same way as in other languages,Switch case also works same.
	for i:=1;i<=5;i++{
		if i%2 == 0{
			fmt.Printf("Even value:%v \n",i)
		}else{
			fmt.Printf("Even value:%v \n",i)
		}
	}
}