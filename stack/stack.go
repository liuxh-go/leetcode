package stack

import "fmt"

type Stack struct {
	Data []int
	L    int
}

// 入栈
func (s *Stack) Push(data int) {
	s.Data = append(s.Data, data)
	s.L++
}

// 出栈
func (s *Stack) Pop() *int {
	if s.L <= 0 {
		return nil
	}

	data := s.Data[s.L-1]
	s.Data = s.Data[:s.L-1]
	s.L--

	return &data
}

// 获取栈顶
func (s *Stack) Top() *int {
	if s.L <= 0 {
		return nil
	}

	return &s.Data[s.L-1]
}

// 栈是否为空
func (s *Stack) Sum() int {
	result := 0
	for _, digit := range s.Data {
		result += digit
	}

	return result
}

// 打印
func (s *Stack) Print() {
	fmt.Println(s.Data)
}

// 新建栈对象
func NewStack() *Stack {
	return &Stack{
		Data: make([]int, 0),
		L:    0,
	}
}
