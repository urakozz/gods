package rbt
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	bst := &RBT{}

	bst.Insert(50)
	assert.Equal(t, int64(50), bst.root.Value)

	bst.Insert(51)
	// >   51 red
	// > 50 black
	// Inserted to the right
	// Rotate left 50 (root) : 50.right.isRed(), 50.Left == nil [cond #1]
	// > 51 - black
    // >   50 - red
	assert.Equal(t, int64(51), bst.root.Value)
	assert.Equal(t, int64(50), bst.root.Left.Value)

	bst.Insert(25)
	// > 51 - black
	// >   50 - red
	// >     25 - red
	// Rotate right 51 : 51.left.isRed() && 51.left.left.isRed() [cond #2]
	// >   51 - red
    // > 50 - black
	// >   25 - red
	// Flip colors 50 : 50.left.isRed() && 50.right.isRed() [cond #3]
	// >   51 - black
	// > 50 - black
	// >   25 - black
	assert.Equal(t, int64(50), bst.root.Value)
	assert.Equal(t, int64(51), bst.root.Right.Value)
	assert.Equal(t, int64(25), bst.root.Left.Value)


	bst.Insert(100)
	// >     100 - red
	// >   51 - black
	// > 50 - black
	// >   25 - black
	// Rotate left 51: 51.right.isRed(), 51.Left == nil [cond #1]
	// >   100 - black
	// >     51 - red
	// > 50 - black
	// >   25 - black
	assert.Equal(t, int64(50), bst.root.Value)
	assert.Equal(t, int64(100), bst.root.Right.Value)
	assert.Equal(t, int64(51), bst.root.Right.Left.Value)
	assert.Equal(t, int64(25), bst.root.Left.Value)

	//bst.Traverse()




}
