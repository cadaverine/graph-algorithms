package stack

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	stack := Stack{}

	stack.Push('a')
	stack.Push(10)
	stack.Push(10.1)
	stack.Push(nil)
	stack.Push("str")

	result, error := stack.Pop()
	if result != "str" || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != nil || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != 10.1 || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != 10 || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != 'a' || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != 0 || error == nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}
}
