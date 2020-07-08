package stack 

// The type for the stack
type Stack struct {
	i    int
	data [10]int
}

// Pushes a number onto the stack
func (s *stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

// Pops a number from the stack
func (s *stack) Pop() int {
	s.i--
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}