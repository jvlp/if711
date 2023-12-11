package main

import (
	"math/rand"
	"sync"
)

type Node struct {
	Key      int
	Adjacent []*Node
}

func newNode(k int) *Node {
	return &Node{Key: k}
}

func (n *Node) AddEdge(node *Node) {
	n.Adjacent = append(n.Adjacent, node)
}

type Graph struct {
	Vertices []*Node
}

func (g *Graph) AddNode(key int) *Node {
	node := newNode(key)
	g.Vertices = append(g.Vertices, node)
	return node
}

func NewGraph(adj [][]int) *Graph {
	g := &Graph{}
	for i, n := range adj {
		node := g.AddNode(i)
		for _, a := range n {
			node.AddEdge(newNode(a))
		}
	}
	return g
}

func NewRandomGraph(n, e int) *Graph {
	r := rand.New(rand.NewSource(42))
	g := &Graph{}
	for i := 0; i < n; i++ {
		g.AddNode(i)
	}
	for i := 0; i < e; i++ {
		u := r.Intn(n)
		v := r.Intn(n)
		for u == v {
			v = r.Intn(v)
		}
		g.Vertices[u].AddEdge(g.Vertices[v])
	}
	return g
}

func (g *Graph) Print() {
	for _, node := range g.Vertices {
		print(node.Key, " -> ")
		for _, n := range node.Adjacent {
			print(n.Key, " ")
		}
		print("\n")
	}
}

func (g *Graph) Bfs(s int) {
	d := make([]int, len(g.Vertices))
	for i := range d {
		d[i] = -1
	}
	d[s] = 0
	level := 1

	frontier := []*Node{g.Vertices[s]}
	for len(frontier) > 0 {
		var next []*Node
		for _, u := range frontier {
			for _, v := range u.Adjacent {
				if d[v.Key] == -1 {
					d[v.Key] = level
					next = append(next, v)
				}
			}
		}
		frontier = next
		level++
	}
}

func (g *Graph) ParallelBfs1(s int) {
	d := NewSafeArray(len(g.Vertices))
	for i := range d.array {
		d.array[i] = -1
	}
	d.array[s] = 0
	level := 1

	frontier := []*Node{g.Vertices[s]}
	var mutex sync.Mutex
	wg := sync.WaitGroup{}
	for len(frontier) > 0 {
		//fmt.Printf("%d: %d\n", level, len(frontier))
		nextFrontier := make([]*Node, 0)
		wg.Add(len(frontier))
		for _, u := range frontier {
			go func(u *Node) {
				defer wg.Done()
				next := make([]*Node, 0)
				for _, v := range u.Adjacent {
					if d.Read(v.Key) == -1 {
						d.Write(v.Key, level)
						next = append(next, v)
					}
				}
				mutex.Lock()
				nextFrontier = append(nextFrontier, next...)
				mutex.Unlock()
			}(u)
		}

		wg.Wait()

		frontier = nextFrontier
		level++
	}
}

func (g *Graph) ParallelBfs2(s int) {
	batchSize := 500
	d := NewSafeArray(len(g.Vertices))
	for i := range d.array {
		d.array[i] = -1
	}
	d.array[s] = 0
	level := 1

	frontier := []*Node{g.Vertices[s]}
	var mutex sync.Mutex
	for len(frontier) > 0 {
		//fmt.Printf("%d: %d\n", level, len(frontier))
		var nextFrontier []*Node
		var wg sync.WaitGroup

		for i := 0; i < len(frontier); i += batchSize {
			end := i + batchSize
			if end > len(frontier) {
				end = len(frontier)
			}

			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				next := make([]*Node, 0)

				for j := start; j < end; j++ {
					u := frontier[j]
					for _, v := range u.Adjacent {
						if d.Read(v.Key) == -1 {
							d.Write(v.Key, level)
							next = append(next, v)
						}
					}
				}

				mutex.Lock()
				nextFrontier = append(nextFrontier, next...)
				mutex.Unlock()
			}(i, end)
		}

		wg.Wait()

		frontier = nextFrontier
		level++
	}
}
