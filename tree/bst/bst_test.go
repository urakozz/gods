package bst
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	bst := &BST{}

	bst.Insert(50)
	assert.Equal(t, int64(50), bst.root.Value)

	bst.Insert(25)
	assert.Equal(t, int64(25), bst.root.Left.Value)

	bst.Insert(100)
	assert.Equal(t, int64(100), bst.root.Right.Value)

	bst.Insert(125)
	assert.Equal(t, int64(125), bst.root.Right.Right.Value)

	bst.Insert(75)
	assert.Equal(t, int64(75), bst.root.Right.Left.Value)

	bst.Insert(40)
	assert.Equal(t, int64(40), bst.root.Left.Right.Value)

	bst.Insert(12)
	assert.Equal(t, int64(12), bst.root.Left.Left.Value)

	bst.Insert(1)
	assert.Equal(t, int64(1), bst.root.Left.Left.Left.Value)

	bst.Insert(2)
	assert.Equal(t, int64(2), bst.root.Left.Left.Left.Right.Value)

	bst.Insert(3)
	assert.Equal(t, int64(3), bst.root.Left.Left.Left.Right.Right.Value)

	bst.Insert(5)
	assert.Equal(t, int64(5), bst.root.Left.Left.Left.Right.Right.Right.Value)

	bst.Insert(4)
	assert.Equal(t, int64(4), bst.root.Left.Left.Left.Right.Right.Right.Left.Value)

//	bst.Traverse()

	bst.Delete(50)
	assert.Equal(t, int64(75),  bst.root.Value)
	assert.Equal(t, int64(100), bst.root.Right.Value)
	assert.Equal(t, int64(125), bst.root.Right.Right.Value)
	assert.True(t,  nil == bst.root.Right.Left)
	assert.Equal(t, int64(25),  bst.root.Left.Value)


//	bst.Traverse()

	bst.Delete(5)
//	bst.Traverse()

	bst.Delete(4)
//	bst.Traverse()

	bst.Reverse()
//	bst.Traverse()

	bst.Reverse()
//	bst.Traverse()

}
