package main

import "fmt"

type math struct{}

func (m *math) add(data ...int) int {
	sum := 0
	for _, val := range data {
		sum += val
	}
	return sum
}

func main() {
	mt := &math{}
	ans1 := mt.add(2, 3)
	ans2 := mt.add(2, 3, 4)
	ans3 := mt.add(2, 3, 4, 5)

	fmt.Println(ans1, " ", ans2, " ", ans3)
}
