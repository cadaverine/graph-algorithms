package list

import "fmt"

type Node struct {
	Data interface{}
	next *Node
}

// List - реализация односвязного списка
type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

// AddData - добавить данные в конец списка
func (list *List) AddData(data interface{}) {
	list.AddNode(&Node{data, nil})
}

// AddNode - добавить элемент в конец списка
func (list *List) AddNode(node *Node) {
	if list.Tail == nil {
		list.Tail = node
		list.Head = list.Tail
	} else {
		list.Tail.next = node
		list.Tail = node
	}
	list.Length++
}

// String - метод для вывода списка
func (list *List) String() string {
	s := ""
	node := list.Head
	for {
		s += fmt.Sprintf("%s", node.Data)

		node = node.next

		if node != nil {
			s += ", "
		} else {
			break
		}
	}

	return s
}
