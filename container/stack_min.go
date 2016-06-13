package container
type MinStack struct{
    list *StackInt
    mins *StackInt
    min int
}

func NewMinStack()*MinStack{
    return &MinStack{list:&StackInt{}, mins:&StackInt{}}
}

func (s *MinStack) Push(i int) {
    s.list.Push(i)
    if s.list.IsEmpty() || i < s.min {
        s.min = i
    }
    s.mins.Push(s.min)
}

func (s *MinStack) Pop() int {
    i, _ := s.list.Pop()
    s.mins.Pop()
    s.min = s.mins.GetLast()
    return i
}

func (s *MinStack) GetMin() int{
    return s.min
}


