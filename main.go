package main

import "fmt"

func Calculate(i int)(result int){
	result = i + 2
	return result
}

func main(){
	fmt.Println("Testing calculate for 2");
	result := Calculate(2)
	fmt.Println(result)
}