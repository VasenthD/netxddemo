package main

import "fmt"

type node struct {
	prev *node
	data int
	next *node
}

type list struct {
	head *node
	tail *node
}

func createnode(n int) *node {
	return &node{
		prev: nil,
		data: n,
		next: nil,
	}
}

func addfirst(n int, head **node) {
	var newnode *node = createnode(n)
	if *head == nil {
		*head = newnode
	} else {
		newnode.next = *head
		(*head).prev = newnode
		*head = newnode
	}
}

func addlast(n int, head **node) {
	var newnode *node = createnode(n)
	if *head == nil {
		*head = newnode
	} else {
		var curr *node = *head
		for curr != nil {
			if curr.next != nil {
				curr = curr.next
			} else {
				break
			}

		}
		curr.next = newnode
		newnode.prev = curr

	}
}

func display(head *node) {
	var current *node = head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func deletefirst(head **node) {
	if *head == nil {
		return
	}
	*head = (*head).next
	(*head).prev = nil

}

func deletelast(head **node) {

	if *head == nil {
		return
	} else {
		var curr *node = *head
		for curr.next != nil {
			curr = curr.next
			if curr.next == nil {
				break
			}
		}
		curr = curr.prev
		curr.next = nil

	}
}

func main() {

	var head *node = nil

	addfirst(5, &head)
	addfirst(4, &head)
	addfirst(3, &head)
	addfirst(2, &head)
	addfirst(1, &head)
	addlast(6, &head)
	addlast(7, &head)
	addlast(8, &head)
	display(head)
	fmt.Println("im done")
	deletefirst(&head)
	deletefirst(&head)
	deletelast(&head)
	display(head)

}
