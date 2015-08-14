package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewWishlist(t *testing.T) {

	str := CString("a")
	cmp := str.CompareTo(CString("b"))
	assert.Equal(t, -1, cmp)

}