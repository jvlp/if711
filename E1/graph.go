package main

import (
	"math/rand"
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
	g := &Graph{}
	for i := 0; i < n; i++ {
		g.AddNode(i)
	}
	for i := 0; i < e; i++ {
		u := rand.Intn(n)
		v := rand.Intn(n)
		for u == v {
			v = rand.Intn(v)
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
