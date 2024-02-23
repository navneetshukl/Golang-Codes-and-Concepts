package main

import "fmt"

type math struct{}

func (m *math) add(a, b int) int {
	return a + b
}

func (m *math) add(a, b, c int) int {
	return a + b + c
}

func main() {
	mt := &math{}
	ans := mt.add(2, 3)
	ans1 := mt.add(2, 3, 4)
	fmt.Println(ans, " ", ans1)
}
