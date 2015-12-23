// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package graph

type UndirectedGraph struct {
	v int
	e int
	adj [][]int
}

func NewUndirectedGraph(v int) *UndirectedGraph {
	g := &UndirectedGraph{v:v}
	g.adj = make([][]int, g.v)
	for i, _ := range g.adj {
		g.adj[i] = []int{}
	}
	return g
}

func (g *UndirectedGraph) V() int {
	return g.v
}
func (g *UndirectedGraph) E() int {
	return g.e
}
func (g *UndirectedGraph) Adj(v int) []int {
	if g.adj != nil && len(g.adj)>v {
		return g.adj[v]
	}
	return []int{}
}
func (g *UndirectedGraph) AddEdge(v,w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e ++
}

var _ IGraph = (*UndirectedGraph)(nil)

