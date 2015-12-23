// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package container
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStackInt(t *testing.T) {

	s := &StackInt{}
	assert.Nil(t, s.stack)

	s.Push(1)
	assert.Equal(t, []int{1}, s.stack)

	s.Push(2)
	assert.Equal(t, []int{1,2}, s.stack)

	v, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, v, 2)
	assert.Equal(t, []int{1}, s.stack)
}
