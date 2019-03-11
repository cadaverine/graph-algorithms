package list

type Node struct {
	data interface{}
	next *Node
}

// List - реализация односвязного списка
type List struct {
	head   *Node
	tail   *Node
	length int
}

// AddData - добавить данные в конец списка
func (list *List) AddData(data interface{}) {
	list.AddNode(&Node{data, nil})
}

// AddNode - добавить элемент в конец списка
func (list *List) AddNode(node *Node) {
	if list.tail == nil {
		list.tail = node
		list.head = list.tail
	} else {
		list.tail.next = node
		list.tail = node
	}
	list.length++
}
