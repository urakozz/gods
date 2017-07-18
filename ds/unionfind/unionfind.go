package unionfind

import "sync"

type UnionFind struct {
	sync.RWMutex
	nodes map[uint64]*Node
}

type Node struct {
	sync.RWMutex
	Parent *Node
	Rank   uint64
	ID     uint64
}

func (n *Node) Init() {
	if n == nil {
		n = &Node{}
	}
	if n.Parent == nil {
		n.Parent = n
	}
}

func (n *Node) FindParent() *Node {
	if n == n.Parent {
		return n
	}
	return n.Parent.FindParent()
}

func (x *Node) Union(y *Node) {
	xRoot, yRoot := x.FindParent(), y.FindParent()
	if xRoot == yRoot {
		return
	}
	xRoot.Lock()
	yRoot.Lock()
	defer xRoot.Unlock()
	defer yRoot.Unlock()
	if xRoot.Rank < yRoot.Rank {
		xRoot.Parent = yRoot
	} else if xRoot.Rank > yRoot.Rank {
		yRoot.Parent = xRoot
	} else {
		yRoot.Parent = xRoot
		yRoot.Rank++
	}
}

func NewUnionFind() *UnionFind {
	return &UnionFind{nodes: make(map[uint64]*Node)}
}

func (u *UnionFind) MakeSet(n *Node) *Node {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.nodes[n.ID]; !ok {
		n.Init()
		u.nodes[n.ID] = n
	}
	return n
}

func (u *UnionFind) Find(n *Node) *Node {
	return n.FindParent()
}

func (u *UnionFind) Union(x, y *Node) {
	x.Union(y)
}
