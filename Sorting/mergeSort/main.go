package main

import "fmt"

func merge(st, mid, en int, arr []int) {
	i, j := st, mid+1

	newArray := []int{}

	for i <= mid && j <= en {

		if arr[i] <= arr[j] {
			newArray = append(newArray, arr[i])
			i++
		} else {
			newArray = append(newArray, arr[j])
			j++
		}

	}

	for i <= mid {
		newArray = append(newArray, arr[i])
		i++
	}

	for j <= en {
		newArray = append(newArray, arr[j])
		j++
	}

	n := len(newArray)
	k := st

	for i := 0; i < n; i++ {
		arr[k] = newArray[i]
		k++
	}
}

func mergeSort(st, en int, arr []int) {
	if st < en {
		mid := (st + en) / 2

		mergeSort(st, mid, arr)
		mergeSort(mid+1, en, arr)
		merge(st, mid, en, arr)
	}
}

func main() {
	arr := []int{2, 2, 8, 2, 8, 5, 2, 14, 3, 4, 6, 5, 3, 2, 4, 7, 8, 62, 8, 3, 4, 5}
	n := len(arr)
	mergeSort(0, n-1, arr)

	fmt.Println("The result of MergeSort is ")
	fmt.Println(arr)
}
