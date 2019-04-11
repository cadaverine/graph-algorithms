package heap

import "fmt"

// Item - элемент двоичной кучи
type Item struct {
	priority int
	index    int
	data     interface{}
}

func (item *Item) String() string {
	return fmt.Sprint("{p: ", item.priority, ", d: ", item.data, "}")
}

// OrderType - тип частичного порядка
type OrderType int

const (
	// Minimum - куча сортирована по минимальному приоритету
	Minimum OrderType = iota
	// Maximum - куча сортирована по максимальному приоритету
	Maximum
)

// Heap - структура данных двоичная куча
type Heap struct {
	orderType OrderType
	data      []*Item
}

// Init - конструктор двоичной кучи
func Init(orderType OrderType) *Heap {
	return &Heap{orderType, make([]*Item, 0)}
}

// Size - размер двоичной кучи
func (heap *Heap) Size() int {
	return len(heap.data)
}

// Enqueue - добавить в очередь
func (heap *Heap) Enqueue(object interface{}, priority int) {
	item := &Item{priority, heap.Size(), object}
	heap.data = append(heap.data, item)

	heap.fixUp(heap.Size() - 1)
}

// Dequeue - удалить из очереди
func (heap *Heap) Dequeue() (interface{}, int) {
	heap.swap(0, heap.Size()-1)
	item := heap.pop()

	if heap.Size() != 0 {
		heap.fixDown(0)
	}

	return item.data, item.priority
}

func (heap *Heap) String() string {
	temp := ""

	for _, item := range heap.data {
		temp += fmt.Sprintln(item)
	}

	return temp
}

// ***************
// PRIVATE METHODS
// ***************

func (heap *Heap) pop() *Item {
	item := heap.data[heap.Size()-1]
	heap.data = heap.data[:heap.Size()-1]
	return item
}

func (heap *Heap) swap(i, j int) {
	heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
	heap.data[i].index, heap.data[j].index = heap.data[j].index, heap.data[i].index
}

func (heap *Heap) compare(first, second *Item) bool {
	if heap.orderType == Maximum {
		return first.priority > second.priority
	}

	return first.priority < second.priority
}

func (heap *Heap) getActualChild(i int) *Item {
	leftChildIndex := i*2 + 1
	rightChildIndex := i*2 + 2

	if heap.Size() > rightChildIndex {
		leftChild := heap.data[leftChildIndex]
		rightChild := heap.data[rightChildIndex]

		if heap.compare(leftChild, rightChild) {
			return leftChild
		}
		return rightChild
	} else if heap.Size() > leftChildIndex {
		return heap.data[leftChildIndex]
	} else {
		return nil
	}
}

func (heap *Heap) fixUp(i int) {
	child := heap.data[i]
	parent := heap.data[i/2]

	if heap.compare(child, parent) {
		heap.swap(i, i/2)
		heap.fixUp(i / 2)
	}
}

func (heap *Heap) fixDown(index int) {
	parent := heap.data[index]
	child := heap.getActualChild(index)

	if child != nil && heap.compare(child, parent) {
		heap.swap(parent.index, child.index)
		heap.fixDown(parent.index)
	}
}
