package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)

	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	visited := map[int]bool{}
	var dfs func(int) bool
	dfs = func(node int) bool {
		if node == destination {
			return true
		}

		if visited[node] {
			return false
		}

		visited[node] = true

		for i := range graph[node] {
			if dfs(graph[node][i]) {
				return true
			}
		}

		return false
	}

	return dfs(source)
}
