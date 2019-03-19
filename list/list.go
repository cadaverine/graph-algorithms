package list

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
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
		list.Tail.Next = node
		list.Tail = node
	}
	list.Length++
}

// String - метод для вывода списка (Stringer interface)
func (list *List) String() (out string) {
	node := list.Head
	for node != nil {
		out += fmt.Sprint(node.Data)

		if node.Next != nil {
			out += ", "
			node = node.Next
		} else {
			break
		}
	}

	return out
}
