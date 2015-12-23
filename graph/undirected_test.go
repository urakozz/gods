// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package graph
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUndirectedNew(t *testing.T){
	g := NewUndirectedGraph(3)
	assert.Equal(t, 3, g.V())
	assert.Equal(t, 0, g.E())
	assert.Equal(t, []int{}, g.adj[0])
	assert.Equal(t, []int{}, g.adj[1])
	assert.Equal(t, []int{}, g.adj[2])
}

func TestUndirectedAddEdge(t *testing.T){
	g := NewUndirectedGraph(3)
	g.AddEdge(0,1)
	g.AddEdge(2,1)
	assert.Equal(t, 3, g.V())
	assert.Equal(t, 2, g.E())
	assert.Equal(t, []int{1}, g.adj[0])
}

func TestUndirectedAdj(t *testing.T){
	g := &UndirectedGraph{}
	slice := g.Adj(1)
	assert.Equal(t, []int{}, slice)
}