package main

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{val: -1, next: nil}
}

func (head *MyLinkedList) Get(index int) int {
	cur := head
	for j := 0; cur != nil; j++ {
		if j == index {
			return cur.val
		}
		cur = cur.next
	}
	return -1
}

func (head *MyLinkedList) AddAtHead(val int) {
	if head.val == -1 {
		head.val = val
		return
	}
	head.next = &MyLinkedList{val: head.val, next: head.next}
	head.val = val
}

func (head *MyLinkedList) AddAtTail(val int) {
	if head.val == -1 {
		head.val = val
		return
	}
	cur := head
	for j := 0; ; j++ {
		if cur.next == nil {
			cur.next = &MyLinkedList{val: val, next: nil}
			return
		}
		cur = cur.next
	}
}

func (head *MyLinkedList) AddAtIndex(index int, val int) {
	cur := head
	if index == 0 {
		head.AddAtHead(val)
	}
	for j := 0; cur != nil; j++ {
		if j == index-1 {
			cur.next = &MyLinkedList{val: val, next: cur.next}
			return
		}
		cur = cur.next
	}
}

func (head *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 && head.next != nil {
		head.val = head.next.val
		head.next = head.next.next
		return
	}
	if index == 0 && head.next == nil {
		head.val = -1
		head.next = nil
		return
	}
	cur := head
	for j := 0; cur != nil; j++ {
		if j == index-1 && cur.next != nil {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}
