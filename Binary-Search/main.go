package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, k int) bool {
	sort.Ints(arr)
	st, en := 0, len(arr)-1

	for st <= en {
		mid := (st + en) / 2

		if arr[mid] == k {
			return true
		} else if arr[mid] < k {
			st = mid + 1
		} else {
			en = mid - 1
		}
	}
	return false

}

func main() {
	arr := []int{5, 6, 7, 5, 7, 4, 6, 7, 5, 6, 8, 7, 7, 5, 5, 63, 3, 4, 1, 2, 3, 3, 4, 5, 87, 0, 9, 7, 9, 57}
	var k int

	fmt.Println("Please enter the value to be searched")

	fmt.Scanln(&k)
	ans := binarySearch(arr, k)

	if ans {
		fmt.Println(fmt.Sprintf("The value %d is present ", k))
	} else {
		fmt.Println(fmt.Sprintf("The value %d is not present ", k))
	}

}
