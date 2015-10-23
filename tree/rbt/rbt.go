// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package rbt

import (
	"fmt"
	"strings"
	"errors"
)

type Node struct {
	Value int64
	Left  *Node
	Right *Node
	Color bool
}

func (n *Node) isRed() bool {
	return n.Color == RED
}

func (n *Node) rotateLeft() *Node {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.Color = n.Color
	n.Color = RED

	//fmt.Printf("Rotate left %d, return %+v, left %+v, right %+v \n", n.Value, x, x.Left, x.Right )
	return x
}

func (n *Node) rotateRight() *Node {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.Color = n.Color
	n.Color = RED

	//fmt.Printf("Rotate right %d, return %+v \n", n.Value, x)
	return x
}

func (n *Node) flipColors() {
	n.Color = RED
	n.Right.Color, n.Left.Color = BLACK, BLACK
	//fmt.Printf("Flip %d \n", n.Value)
}

const RED = true
const BLACK = false

type RBT struct {
	root *Node
}

func NewRedBlackTree() *RBT {
	return &RBT{}
}


func (b *RBT) Insert(value int64) {

	b.root = insert(b.root, value)
	b.root.Color = BLACK
}

func (b *RBT) Min() (int64, error) {
	if b.root != nil {
		return b.root.FindMin().Value, nil
	}
	return 0, errors.New("BST is empty")
}

func (b *RBT) Max() (int64, error) {
	if b.root != nil {
		return b.root.FindMax().Value, nil
	}
	return 0, errors.New("BST is empty")
}

func (b *RBT) Reverse() {
	b.root.Reverse()
}

func (b *RBT) Traverse() {
	b.root.Traverse(0)
}

//func (b *BST) Delete(key int64) {
//	if b.root != nil {
//		b.root = deleteNode(b.root, key)
//	}
//}

func (n *Node) Traverse(level int) {
	if n.Right != nil { n.Right.Traverse(level + 1) }
	fmt.Printf("> %s%d - %+v\n", strings.Repeat("\t", level), n.Value, n.Color)
	if n.Left != nil { n.Left.Traverse(level + 1) }
}
func (n *Node) Reverse() {
	n.Right, n.Left = n.Left, n.Right
	if n.Right != nil { n.Right.Reverse() }
	if n.Left != nil { n.Left.Reverse() }
}

func (n *Node) FindMin() *Node {
	if n.Left != nil {
		return n.Left.FindMin()
	}
	return n
}

func (n *Node) FindMax() *Node {
	if n.Right != nil {
		return n.Right.FindMax()
	}
	return n
}

func insert(n *Node, value int64) *Node {
	if n == nil {
		n = &Node{Value:value, Color:RED}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value:value, Color:RED}
		}
		n.Right = insert(n.Right, value)
	} else if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value:value, Color:RED}
		}
		n.Left = insert(n.Left, value)
	}

	if (n.Right != nil && n.Right.isRed()) && (n.Left == nil || !n.Left.isRed()) {
		n = n.rotateLeft()
	}
	if (n.Left != nil && n.Left.isRed()) && (n.Left.Left != nil && n.Left.Left.isRed()) {
		n = n.rotateRight()
	}
	if (n.Left != nil && n.Left.isRed()) && (n.Right != nil && n.Right.isRed()) {
		n.flipColors()
	}

	return n
}
//
//func deleteMax(n *Node) *Node {
//	if n.Right == nil {
//		return n.Left
//	}
//	n.Right = deleteMax(n.Right)
//	return n
//}
//
//func deleteMin(n *Node) *Node {
//	if n.Left == nil {
//		return n.Right
//	}
//	n.Left = deleteMin(n.Left)
//	return n
//}
//
//func deleteNode(n *Node, key int64) *Node {
//	//	fmt.Printf("%+v\n", n)
//	if key < n.Value {
//		n.Left = deleteNode(n.Left, key)
//	} else if key > n.Value {
//		n.Right = deleteNode(n.Right, key)
//	} else {
//		// Remove this node
//		if n.Left == nil && n.Right == nil {
//			return nil
//		}
//		if n.Left == nil {
//			return n.Right
//		}
//		if n.Right == nil {
//			return n.Left
//		}
//		//		n.Traverse(0)
//		t := n
//		n = t.Right.FindMin()
//		n.Right = deleteMin(t.Right)
//		n.Left = t.Left
//
//
//	}
//
//	return n
//}


