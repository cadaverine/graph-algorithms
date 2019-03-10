package stack

import "errors"

// Stack - простая реализация стека
type Stack struct {
	Data []interface{}
}

// Push - поместить элемент в стек
func (stack *Stack) Push(value interface{}) {
	stack.Data = append(stack.Data, value)
}

// Pop - убрать элемент из стека
func (stack *Stack) Pop() (interface{}, error) {
	length := stack.Size()

	if length > 0 {
		last := stack.Data[length-1]
		stack.Data = stack.Data[:length-1]
		return last, nil
	}

	return nil, errors.New("stack is empty")
}

// IsEmpty - проверка стека на наличие элементов
func (stack *Stack) IsEmpty() bool {
	return len(stack.Data) == 0
}

// Size - количество элементов в стеке
func (stack *Stack) Size() int {
	return len(stack.Data)
}
