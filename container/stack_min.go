type MinStack struct{
    list *Stack
    mins *Stack
    min int
}

func NewMinStack()*MinStack{
    return &MinStack{list:NewStack(), mins:NewStack()}
}

func (s *MinStack) Push(i int) {
    s.list.Push(i)
    if s.list.IsEmpty() || i < s.min {
        s.min = i
    }
    s.mins.Push(s.min)
}

func (s *MinStack) Pop() int {
    i := s.list.Pop()
    s.mins.Pop()
    s.min = s.mins.GetLast()
    return i
}

func (s *MinStack) GetMin() int{
    return s.min
} 


type Stack struct{
    list []int
}
func NewStack() * Stack{
    return &Stack{list:make([]int)}
}
func (s *MinStack) Push(i int) {
    s.list = append(s.list, i)
}
func (s *MinStack) Pop() int {
    ans := s.GetLast()
    s.list = s.list[0:len(s.list)-2]
    return ans
}
func (s *MinStack) IsEmpty() bool {
    return 0 == len(s.list)
}
func (s *MinStack) GetLast() int {
    return s.list[len(s.list)-1]
}
