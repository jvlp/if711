package main

const (
	NVertices = 1_000_000
	NEdges    = 10_000_000
)

func main() {
	//adj := [][]int{{3, 2}, {0, 2}, {0, 1, 3}, {2}}
	//g := NewGraph(adj)

	g := NewRandomGraph(NVertices, NEdges)
	// g.Print()
	// printDistances(g.Bfs(0))
	//start := time.Now()
	// g.Bfs(0)
	//elapsed := time.Since(start)
	//log.Printf("Bfs took %s", elapsed)
	//start = time.Now()
	g.ParallelBfs(0)
	//elapsed = time.Since(start)
	//log.Printf("ParallelBfs took %s", elapsed)
}
