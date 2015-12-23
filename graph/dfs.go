// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package graph
import "container/list"

type DeepFirstSearch struct {
	marked []bool
	count int
}

func NewDFS(g IGraph, s int) *DeepFirstSearch{
	d := &DeepFirstSearch{}
	d.marked = make([]bool, d.count)
	d.dfs(g, s)
	list.New()
	return d
}


func (d *DeepFirstSearch) dfs(g IGraph, v int) {
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

