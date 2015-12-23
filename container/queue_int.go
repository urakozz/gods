// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package container

type QueueInt struct {
	in *StackInt
	out *StackInt
}

func (q *QueueInt) Enqueue(v int) {
	if q.in == nil {
		q.in = &StackInt{}
	}
	if q.out == nil {
		q.out = &StackInt{}
	}
	q.in.Push(v)
}

func (q *QueueInt) Dequeue() (int, error){
	if !q.out.IsEmpty(){
		return q.out.Pop()
	}

	if q.out.IsEmpty() {
		for !q.in.IsEmpty(){
			tmp, _ := q.in.Pop()
			q.out.Push(tmp)
		}
	}
	return q.out.Pop()
}


