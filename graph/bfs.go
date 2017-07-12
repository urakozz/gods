// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package graph
import "github.com/urakozz/gods/container"


type BreadthFirstSearch struct {
	q container.QueueInt
	marked []bool
	count int
}

func NewBFS(g IIntGraph, s int) *BreadthFirstSearch{
	d := &BreadthFirstSearch{}
	d.marked = make([]bool, d.count)

	d.bfs(g, s)
	return d
}

func (d *BreadthFirstSearch) bfs(g IIntGraph, s int) {
	d.q.Enqueue(s)
	d.marked[s] = true

	for !d.q.IsEmpty() {
		v, _ := d.q.Dequeue()
		for _, a := range g.Adj(v) {
			if ! d.marked[a]{
				d.marked[a] = true
				d.count++
				d.q.Enqueue(a)
			}
		}
	}

}

func (d *BreadthFirstSearch) Count() int {
	return d.count
}
func (d *BreadthFirstSearch) Marked(v int) bool {
	return d.marked[v]
}

