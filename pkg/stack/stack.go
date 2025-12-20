package stack

type Stack struct {
	elements []int
}

func New() *Stack {
	return &Stack{elements: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.elements = append(s.elements, v)
}

func (s *Stack) Pop() int {
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func (s *Stack) Peek() int {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Empty() bool {
	return len(s.elements) == 0
}
