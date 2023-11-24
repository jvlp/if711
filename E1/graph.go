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

func NewRandomGraph(v, e int) *Graph {
	g := &Graph{}
	for i := 0; i < v; i++ {
		g.AddNode(i)
	}
	for i := 0; i < e; i++ {
		u := rand.Intn(v)
		v := rand.Intn(v)
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
		println()
	}
}

func (g *Graph) Bfs(s int) {
	queue := []int{s}
	visited := make([]bool, len(g.Vertices))
	visited[s] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		print(u, ", ")
		for _, v := range g.Vertices[u].Adjacent {
			if !visited[v.Key] {
				visited[v.Key] = true
				queue = append(queue, v.Key)
			}
		}
	}
}
