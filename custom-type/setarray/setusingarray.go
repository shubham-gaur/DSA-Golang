package setarray

import (
	"fmt"
	"reflect"
)

type set[T any] []T

type SetServicer interface {
	Add(ele any)
	Remove(ele any)
	Length() int
	IsEmpty() bool
	Contains(ele any) bool
	IndexOf(ele any) int
	Print()
}

// Add: It adds an element of type any within the set.
func (s *set[T]) Add(ele T) {
	for _, val := range *s {
		if reflect.ValueOf(val) == reflect.ValueOf(ele) {
			return
		}
	}
	*s = append(*s, ele)
}

// Remove: It removes an element of type any within the set.
func (s *set[T]) Remove(ele T) {
	temp := []T{}
	for _, val := range *s {
		if reflect.ValueOf(val) != reflect.ValueOf(ele) {
			temp = append(temp, val)
		}
	}
	*s = temp
}

// Length: This computes the total length of array. Returns the integer value.
func (s *set[T]) Length() int {
	return len(*s)
}

// IsEmpty: Returns true if length is grether than 0.
func (s *set[T]) IsEmpty() bool {
	if len(*s) > 0 {
		return false
	} else {
		return true
	}
}

// Contains: Returns true if an element is present in the set.
func (s *set[T]) Contains(ele T) bool {
	for _, val := range *s {
		if reflect.ValueOf(val) == reflect.ValueOf(ele) {
			return true
		}
	}
	return false
}

// IndexOf: Returns the index of the element in the set.
func (s *set[T]) IndexOf(ele T) int {
	for key, val := range *s {
		if reflect.ValueOf(val) == reflect.ValueOf(ele) {
			return key
		}
	}
	return -1
}

// Print: Responsible to print all the elements of the set.
func (s *set[T]) Print() {
	fmt.Println(*s)
}

func NewSet() SetServicer {
	var s set[any]
	return &s
}

func RunSetInArray() {
	s := NewSet()
	s.Add(4)
	s.Add(4)
	s.Add(5)
	s.Add(6)
	s.Add(7)
	s.Add(8)
	s.Add(8)
	fmt.Println("Length    : ", s.Length())
	s.Print()
	s.Remove(4)
	fmt.Println("Length    : ", s.Length())
	s.Print()
	s.Remove(5)
	fmt.Println("Length    : ", s.Length())
	s.Print()
	fmt.Println("Contains 7: ", s.Contains(7))
	fmt.Println("Index of 7: ", s.IndexOf(7))
}
