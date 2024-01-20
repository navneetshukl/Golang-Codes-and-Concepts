package main

import "fmt"

// Define the Node structure for doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// Define the DoublyLinkedList structure
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Initialize a new empty doubly linked list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Add a new node at the end of doubly linked list
func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{data: data, prev: nil, next: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
}

// Display the doubly linked list in forward direction
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.head

	for current.next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}

// Display the doubly linked list in a backward direction
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}

	fmt.Println("nil")
}

func main() {

	// Create a new doubly linked list
	dll := NewDoublyLinkedList()

	//Add elements to the doubly linked list

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	// Display the doubly linked list in forward direction
	fmt.Println("Doubly Linked List (Forward):")
	dll.DisplayForward()

	// Display the doubly linked list in backward direction
	fmt.Println("Doubly Linked List (Backward):")
	dll.DisplayBackward()
}
