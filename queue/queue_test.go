package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}

	queue.Enqueue('a')
	queue.Enqueue(10)
	queue.Enqueue(10.1)
	queue.Enqueue(nil)
	queue.Enqueue("str")

	result, error := queue.Dequeue()
	if result != 'a' || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = queue.Dequeue()
	if result != 10 || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = queue.Dequeue()
	if result != 10.1 || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = queue.Dequeue()
	if result != nil || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = queue.Dequeue()
	if result != "str" || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = queue.Dequeue()
	if result != nil || error == nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}
}
