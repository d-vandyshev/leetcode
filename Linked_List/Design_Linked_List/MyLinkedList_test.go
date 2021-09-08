package main

import "testing"

func TestMyLinkedList(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	val := myLinkedList.Get(1)
	if val != 2 {
		t.Error("Test1 expected", 2, "got", val)
	}
	myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3
	val = myLinkedList.Get(1)
	if val != 3 {
		t.Error("Test2 expected", 3, "got", val)
	}
}
