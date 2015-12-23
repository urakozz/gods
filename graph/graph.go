// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package graph

type IGraph interface {
	V() int
	E() int
	Adj(v int) []int
	AddEdge(v,w int)
}

