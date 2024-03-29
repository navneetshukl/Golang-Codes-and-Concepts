package main

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("Hi from say function")
}

type child struct {
	base  //!embedding
	style string
}

func main() {
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "Babu Rao ka style hain",
	}

	child.say()
	fmt.Println("The color is ", child.color)

}
