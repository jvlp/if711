package main

func main() {
	//adj := [][]int{{3, 2}, {0, 2}, {0, 1, 3}, {2}}
	//g := NewGraph(adj)

	g := NewRandomGraph(100, 1000)
	g.Print()
	g.Bfs(0)
}
