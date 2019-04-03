package graph

// Item - элемент двоичной кучи
type Item struct {
	priority int
	index    int
	data     interface{}
}

// OrderType - тип частичного порядка
type OrderType int

const (
	// Minimum - куча сортирована по минимальному приоритету
	Minimum OrderType = iota
	// Maximum - куча сортирована по максимальному приоритету
	Maximum
)

// Heap - структура данных - двоичная куча
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
func (heap *Heap) Dequeue() interface{} {
	heap.swap(0, heap.Size()-1)
	item := heap.pop()
	heap.fixDown(0)

	return item.data
}

func (heap *Heap) push() {}

func (heap *Heap) pop() *Item {
	item := heap.data[heap.Size()-1]
	heap.data = heap.data[:heap.Size()-1]
	return item
}

func (heap *Heap) compare(child, parent *Item) bool {
	if heap.orderType == Maximum {
		return child.priority > parent.priority
	}

	return child.priority < parent.priority
}

func (heap *Heap) swap(i, j int) {
	heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
}

func (heap *Heap) fixUp(i int) {
	// добавить проеверки на существование элементов по указанным индексам
	child := heap.data[i]
	parent := heap.data[i/2]

	// проверить compare на корректность работы с OrderType
	if heap.compare(child, parent) {
		heap.swap(i, i/2)
		heap.fixUp(i / 2)
	}
}

func (heap *Heap) getMaxChild(i int) *Item {
	leftChildIndex := i*2 + 1
	rightChildIndex := i*2 + 2

	if heap.Size() > leftChildIndex {

	} else if heap.Size() > leftChildIndex {
		return heap.data[leftChildIndex]
	} else {
		return nil
	}
}

func (heap *Heap) fixDown() {
	parent := heap.data[0]
	child := heap.getMaxChild()

	if child != nil {
		heap.swap()
	}
	// child := math.Max(heap.data[i*2+1].priority, heap.data[i*2+2].priority) использовать compare для сравнения

	if heap.compare(child, parent) {

	}
}
