package main

import "fmt"

func main(){
	var x int=20

	var ptr *int=&x

	fmt.Println("Given value is ",x)
	fmt.Println("Address of the given value is ",ptr)
	fmt.Println("Original value from pointer is ",*ptr)

}