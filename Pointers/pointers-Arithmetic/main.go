package main

import "fmt"


func main() {
	var x int = 20
	var ptr *int = &x

	fmt.Println("Address of the value x is ", ptr)
	fmt.Println("Value from the pointer is ", *ptr)

	fmt.Println("Before Adding 1, value is ",*ptr)
	fmt.Println("After Adding 1, value is ",*ptr+1)

	fmt.Println("Before Substracting 1,value is ",*ptr)
	fmt.Println("After Substracting 1, value is ",*ptr-1)

}
