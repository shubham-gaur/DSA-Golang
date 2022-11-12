package datastructures

type StackServicer interface {
	Push(ele string)
	Pop()
	Top() string
	GetStack() []string
	IsEmpty() bool
}

type stack struct {
	stackArr []string
}

func InitStack(s string) StackServicer {
	var stackArr []string
	for i := 0; i < len(s); i++ {
		stackArr = append(stackArr, string(s[i]))
	}
	return &stack{
		stackArr: stackArr,
	}
}

func (s *stack) Push(ele string) {
	s.stackArr = append(s.stackArr, ele)
}

func (s *stack) Pop() {
	s.stackArr = s.stackArr[0 : len(s.stackArr)-1]
}

func (s *stack) Top() string {
	if len(s.stackArr) == 0 {
		return ""
	} else {
		return s.stackArr[len(s.stackArr)-1]
	}
}

func (s *stack) GetStack() []string {
	return s.stackArr
}

func (s *stack) IsEmpty() bool {
	return len(s.stackArr) == 0
}
