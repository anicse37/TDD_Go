package generics

type StackOfInts struct {
	values []int
}
type StackOfStrings struct {
	values []string
}
type Stack[T any] struct {
	values []T
}

/*-----------------------------------------------------------*/

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	value := s.values[index]
	s.values = s.values[:index]
	return value, true
}

/*-----------------------------------------------------------*/

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	value := s.values[index]
	s.values = s.values[:index]
	return value, true
}

/*-------------------------------------------------------*/
func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

/*-------------------------------------------------------*/

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var Zero T
		return Zero, false
	}

	index := len(s.values) - 1
	value := s.values[index]
	s.values = s.values[:index]
	return value, true
}

/*-----------------------------------------------------------*/
