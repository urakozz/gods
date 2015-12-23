// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package container
import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestQueueInt(t *testing.T) {
	q := &QueueInt{}
	q.Enqueue(1)
	assert.Equal(t, []int{1}, q.in.stack)
	assert.Equal(t, []int(nil), q.out.stack)

	v, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	assert.Equal(t, []int{}, q.in.stack)
	assert.Equal(t, []int{}, q.out.stack)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, []int{1,2,3}, q.in.stack)
	assert.Equal(t, []int{}, q.out.stack)

	v, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	assert.Equal(t, []int{}, q.in.stack)
	assert.Equal(t, []int{3,2}, q.out.stack)

	q.Enqueue(10)
	assert.Equal(t, []int{10}, q.in.stack)
	assert.Equal(t, []int{3,2}, q.out.stack)

	v, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 2, v)
	assert.Equal(t, []int{10}, q.in.stack)
	assert.Equal(t, []int{3}, q.out.stack)

	v, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 3, v)
	assert.Equal(t, []int{10}, q.in.stack)
	assert.Equal(t, []int{}, q.out.stack)

	v, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 10, v)
	assert.Equal(t, []int{}, q.in.stack)
	assert.Equal(t, []int{}, q.out.stack)

	v, err = q.Dequeue()
	assert.Equal(t, err, errorStackEmpty)
}




