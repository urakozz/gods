// Copyright 2015 Home24 AG. All rights reserved.
// Proprietary license.
package container
import (
	"errors"
	"sync"
)

var errorStackEmpty = errors.New("Stack is empty")

type StackInt struct {
	sync.Mutex
	stack []int
}

func (s *StackInt) Push(v int) {
	if s.stack == nil {
		s.stack = []int{}
	}
	s.Lock()
	defer s.Unlock()
	s.stack = append(s.stack, v)
}

func (s *StackInt) Pop() (int, error) {
	s.Lock()
	defer s.Unlock()
	if s.IsEmpty() {
		return 0, errorStackEmpty
	}
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return last, nil
}

func (s *StackInt) IsEmpty() bool {
	return s.stack == nil || 0 == len(s.stack)
}

