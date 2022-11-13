package datastructures

import (
	"fmt"
	"reflect"
)

type StackServicer interface {
	Push(ele any)
	Pop()
	Top() any
	GetStack() []any
	IsEmpty() bool
}

type stack struct {
	stackArr  []any
	stackType reflect.Type
}

func InitStack[T string | []int](ele T) StackServicer {
	var stackArr []any
	var stackType reflect.Type
	if arr, ok := any(ele).([]int); ok {
		for i := 0; i < len(ele); i++ {
			stackArr = append(stackArr, arr[i])
		}
		stackType = reflect.TypeOf(len(arr))
	} else if arr, ok := any(ele).(string); ok {
		for i := 0; i < len(ele); i++ {
			stackArr = append(stackArr, arr[i])
		}
		stackType = reflect.TypeOf(ele)
	}

	return &stack{
		stackArr:  stackArr,
		stackType: stackType,
	}
}

func (s *stack) Push(ele any) {
	if s.stackType == reflect.TypeOf(ele) {
		s.stackArr = append(s.stackArr, ele)
	} else {
		fmt.Println("Does not support this type")
	}
}

func (s *stack) Pop() {
	s.stackArr = s.stackArr[0 : len(s.stackArr)-1]
}

func (s *stack) Top() any {
	if len(s.stackArr) == 0 {
		return nil
	} else {
		return s.stackArr[len(s.stackArr)-1]
	}
}

func (s *stack) GetStack() []any {
	return s.stackArr
}

func (s *stack) IsEmpty() bool {
	return len(s.stackArr) == 0
}
