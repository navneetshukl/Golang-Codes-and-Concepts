package main

var mp map[int][]int
var color map[int]int

func join(a, b int) {
	mp[a] = append(mp[a], b)
}

//  dfs function is actual logic for checking graph is bipartite using dfs algorithm
func dfs(src int) bool {
	si := len(mp[src])
	for i := 0; i < si; i++ {
		if color[mp[src][i]] == -1 {
			if color[src] == 0 {
				color[mp[src][i]] = 1
			} else {
				color[mp[src][i]] = 0
			}
			ans := dfs(mp[src][i])
			if ans == false {
				return false
			}
		} else {
			if color[mp[src][i]] == color[src] {
				return false
			}
		}
	}
	return true
}

//isBipartiteDFS function will call from main to check the graph is bipartite
func isBipartiteDFS(graph [][]int) bool {

	n := len(graph)
	mp = make(map[int][]int)
	color = make(map[int]int)
	for i := 0; i < n; i++ {
		newlen := len(graph[i])
		for j := 0; j < newlen; j++ {
			join(i, graph[i][j])
		}
	}

	for i := 0; i < n; i++ {
		color[i] = -1
	}
	var ans bool
	ans = true
	for i := 0; i < n; i++ {
		if color[i] == -1 {
			color[i] = 0
			val := dfs(i)
			ans = ans && val
		}
	}
	return ans

}
