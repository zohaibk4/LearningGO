package stack

// Stack is the object
type Stack struct {
	i    int
	data [10]int
}

// Push adds a number onto the stack
func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

// Pop removes a number from the stack
func (s *Stack) Pop() int {
	s.i--
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}
