// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package graph

type DeepFirstSearch struct {
	marked []bool
	count int
}

func NewDFS(g IIntGraph, s int) *DeepFirstSearch{
	d := &DeepFirstSearch{}
	d.marked = make([]bool, d.count)
	d.dfs(g, s)
	return d
}


func (d *DeepFirstSearch) dfs(g IIntGraph, v int) {
	d.count++
	d.marked[v] = true
	for _, a := range g.Adj(v) {
		if !d.marked[a] {
			d.dfs(g, a)
		}
	}
}

func (d *DeepFirstSearch) Count() int {
	return d.count
}
func (d *DeepFirstSearch) Marked(v int) bool {
	return d.marked[v]
}

