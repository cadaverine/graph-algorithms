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

// String - метод для вывода списка (Stringer interface)
func (list *List) String() (out string) {
	node := list.Head
	for node != nil {
		out += fmt.Sprint(node.Data)

		if node.next != nil {
			out += ", "
			node = node.next
		} else {
			break
		}
	}

	return out
}
