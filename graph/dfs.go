package graph

func DFS(g Graph, s int) ([]bool, []int) {
	marked := make([]bool, g.GetNumVertices())
	edgeTo := make([]int, g.GetNumVertices())

	dfs(g, s, marked, edgeTo)

	return marked, edgeTo
}

func dfs(g Graph, v int, marked []bool, edgeTo []int) {
	marked[v] = true
	for _, w := range g.Adj(v) {
		if !marked[w] {
			edgeTo[w] = v
			dfs(g, w, marked, edgeTo)
		}
	}
}

func HasPathTo(v int, marked []bool) bool{
	return marked[v]
}

func PathTo(s, v int, marked []bool, edgeTo []int) []int {
	path := make([]int,0)
	if !HasPathTo(v, marked){
		return path
	}
	for x := v; x != s; x = edgeTo[x] {
		path = append(path[0:1], path[:]...)
		path[0] = x
	}
	return path
}
