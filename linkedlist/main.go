package main

import (
	"fmt"
	"os"
	"strconv"
)

type linkedList struct {
	number int
	next   *linkedList
}

// insert operation
func (node *linkedList) insert(num int) {
	var temp = &linkedList{}
	temp.number = num
	temp.next = nil
	if node == nil {
		node = temp
	} else {
		var p = &linkedList{}
		p = node
		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}

// for display
func (node *linkedList) display() {
	var p = &linkedList{}
	p = node.next
	for p != nil {
		fmt.Printf("%d ->", p.number)
		p = p.next
	}
}

func main() {
	head := new(linkedList)
	var choice string
	for true {
		fmt.Println("Enter your choice")
		fmt.Println("1. Insert value in linkedList")
		fmt.Println("2. Display linkedList")
		fmt.Println("3. Exit..")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter your value for linkedList node")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insert(num)
		case "2":
			head.display()
		default:
			os.Exit(0)
		}
	}
}
