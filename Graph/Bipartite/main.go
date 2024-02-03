package main

import "fmt"

func main() {
	graph := [][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1, 3},
		{0, 2},
	}

	ans := isBipartiteDFS(graph)
	if ans == true {
		fmt.Println("Graph is Bipartite")
	} else {
		fmt.Println("Graph is not Bipartite")
	}

}
