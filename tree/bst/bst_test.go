package bst
import (
	"testing"
)

func TestInsert(t *testing.T) {
	bst := &BST{}
	bst.Insert(50)
	bst.Insert(25)
	bst.Insert(100)
	bst.Insert(125)
	bst.Insert(75)
	bst.Insert(40)
	bst.Insert(12)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(4)

	bst.Traverse()

	bst.Delete(50)
	bst.Traverse()

	bst.Delete(5)
	bst.Traverse()

	bst.Delete(4)
	bst.Traverse()

}
