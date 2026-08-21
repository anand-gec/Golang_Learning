package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList2 struct {
	head   *node
	length int
}

//to store
func (l *linkedList2) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

//display 
func (l linkedList2) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

//delete
func (l *linkedList2) deleteWithValue(value int) {
	if l.length==0{
return
	}

	if l.head.data==value{
		l.head=l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next==nil{
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList2{}
	node1 := &node{data: 48}
	node2 := &node{data: 58}
	node3 := &node{data: 42}
	node4 := &node{data: 45}
	node5 := &node{data: 5}
	node6 := &node{data: 8}
	node7 := &node{data: 9}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)

	mylist.printListData()

	mylist.deleteWithValue(42)
	mylist.printListData()

	mylist.deleteWithValue(100)   //run time error  // after some changes it's ignore
	mylist.printListData()

	mylist.deleteWithValue(7)   //run time error due it is head node // after some changes it's ignore
	mylist.printListData()
	
	emptyList:= linkedList2{} //run time error  // after some changes it's ignore
	emptyList.deleteWithValue(4)
}
