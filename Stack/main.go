package main

import (
	"errors"
	"fmt"
)

type stack []int

type STACK interface {
	Push(k int)
	Top() (int, error)
	Pop() error
}

func (st *stack) Push(k int) {

	arr := []int{}
	arr = append(arr, k)
	l := len(*st)
	for i := 0; i < l; i++ {
		arr = append(arr, (*st)[i])
	}
	*st = (*st)[:0]
	l = len(arr)

	for i := 0; i < l; i++ {
		*st = append(*st, arr[i])
	}

}

func (st *stack) Top() (int, error) {

	if len(*st) <= 0 {
		return 0, errors.New("error in getting the top element")
	} else {
		return (*st)[0], nil
	}

}
func (st *stack) Pop() error {
	l := len(*st)
	if l <= 0 {
		return errors.New("error in popping the element")
	}
	*st = (*st)[1:l]
	return nil
}

func main() {
	var st stack

	st.Push(1)
	st.Push(4)
	st.Push(2)
	st.Push(7)
	st.Push(5)
	val, err := st.Top()
	if err != nil {
		return
	}
	fmt.Println(st)
	fmt.Println("Top element is ", val)
	_ = st.Pop()

	fmt.Println(st)
	val, _ = st.Top()
	fmt.Println("Top element is ", val)
	_ = st.Pop()
	fmt.Println(st)
}
