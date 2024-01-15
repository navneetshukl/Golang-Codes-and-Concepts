package main

import "fmt"

func Partition(st, en int, arr []int) int {
	i, j := st-1, st
	pivot := arr[en]

	for j < en {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		j++
	}
	i++
	arr[i], arr[en] = arr[en], arr[i]
	return i
}

func QuickSort(st, en int, arr []int) {
	if st < en {
		parIdx := Partition(st, en, arr)
		QuickSort(st, parIdx-1, arr)
		QuickSort(parIdx+1, en, arr)
	}

}

func main() {
	arr := []int{5, 4, 7, 8, 7, 7, 2, 1, 4, 6, 8, 64, 6, 5, 8, 7, 3, 1, 4, 16, 5}

	en := len(arr)
	QuickSort(0, en-1, arr)
	fmt.Println("Sorted array is")
	fmt.Println(arr)
}
