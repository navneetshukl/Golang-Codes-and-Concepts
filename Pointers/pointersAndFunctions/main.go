package main

import "fmt"

func Add(ptr *int) {
	*ptr = *ptr + 2
}

func main() {
	var i int = 32

	var ptr *int = &i

	fmt.Println("Value before the Addition is ", *ptr)
	Add(ptr)
	fmt.Println("Value After the Addition is ", *ptr)

}
