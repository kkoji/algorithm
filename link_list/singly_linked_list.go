package link_list

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(Data interface{}) {
	new_node := &Node{Data: Data}
	if l.Head == nil {
		l.Head = new_node
		return
	}

	last_node := l.Head
	for last_node.Next != nil {
		last_node = last_node.Next
	}
	last_node.Next = new_node
}

func (l *LinkedList) Insert(Data interface{}) {
	new_node := &Node{Data: Data}
	new_node.Next = l.Head
	l.Head = new_node
}

func (l *LinkedList) Print() {
	current_node := l.Head
	for current_node != nil {
		fmt.Println(current_node.Data)
		current_node = current_node.Next
	}
}

func (l *LinkedList) Remove(Data interface{}) {
	current_node := l.Head
	if current_node != nil && current_node.Data == Data {
		l.Head = current_node.Next
		current_node = nil
		return
	}

	var previous_node *Node
	for current_node != nil && current_node.Data != Data {
		previous_node = current_node
		current_node = current_node.Next
	}

	if current_node == nil {
		return
	}

	previous_node.Next = current_node.Next
	current_node = nil
}
