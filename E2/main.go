package main

import (
	"fmt"
	"time"
)

const (
	NVertices = 10_000_000
	NEdges    = 100_000_000
)

func main() {
	//adj := [][]int{{3, 2}, {0, 2}, {0, 1, 3}, {2}}
	//g := NewGraph(adj)

	g := NewRandomGraph(NVertices, NEdges)
	//g.Print()
	//printDistances(g.Bfs(0))
	//start := time.Now()
	//g.Bfs(0)
	//elapsed := time.Since(start)
	//log.Printf("Bfs took %s", elapsed)

	start := time.Now()
	g.ParallelBfs1(0)
	elapsed := time.Since(start)
	fmt.Printf("ParallelBfs1 took %s\n", elapsed)

	start = time.Now()
	g.ParallelBfs2(0)
	elapsed = time.Since(start)
	fmt.Printf("ParallelBfs2 took %s\n", elapsed)
}
