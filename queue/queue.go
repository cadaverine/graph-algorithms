package queue

import "errors"

// Queue - простая реализация очереди
type Queue struct {
	Data []interface{}
}

// Enqueue - добавить в очередь
func (queue *Queue) Enqueue(value interface{}) {
	queue.Data = append(queue.Data, value)
}

// Dequeue - убрать из очереди первый элемент
func (queue *Queue) Dequeue() (interface{}, error) {
	length := queue.Size()

	if length > 0 {
		first := queue.Data[0]
		queue.Data = queue.Data[1:length]
		return first, nil
	}

	return nil, errors.New("queue is empty")
}

// IsEmpty - проверка очереди на наличие элементов
func (queue *Queue) isEmpty() bool {
	return len(queue.Data) == 0
}

// Size - количество элементов в очереди
func (queue *Queue) Size() int {
	return len(queue.Data)
}
