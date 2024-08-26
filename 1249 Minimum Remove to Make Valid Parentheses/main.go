package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minRemoveToMakeValid("())()((("))
}

type Stack struct {
	Items []int
}

func (s *Stack) Push(index int) {
	s.Items = append(s.Items, index)
}

func (s *Stack) Pop() (index int) {
	l := len(s.Items)
	index = s.Items[l-1]
	s.Items = s.Items[0 : l-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func minRemoveToMakeValid(s string) string {
	var (
		stack  = &Stack{Items: make([]int, 0)}
		result strings.Builder
	)

	for _, ch := range s {
		switch string(ch) {
		case "(":
			stack.Push(result.Len())
			result.WriteRune(ch)
		case ")":
			if !stack.IsEmpty() {
				_ = stack.Pop()
				result.WriteRune(ch)
			}
		default:
			result.WriteRune(ch)
		}
	}

	temp := result.String()

	for !stack.IsEmpty() {
		index := stack.Pop()
		temp = temp[:index] + temp[index+1:]
	}

	return temp
}
