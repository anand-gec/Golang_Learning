package main

func node struct{
	data int
	next *node
}

func linkedlist3 struct{
	header *node
	length int
}

//insert data
func (l *linkedlist3)prepend(n *node){
	second:=l.head
	l.head =n
	l.head.next=second
	l.length++
}

//to get data or show 
func (l linkedlist3)printlistdata(){
	toprint:=l.head
	for length !=0{
		fmt.Printf("%d"toprint.data)
		toprint=toprint.next
		l.length--
	}
	fmt.Printf("\n")
}

//to delete data
func (l *linkedlist3)deletewithvalue(value int){
previoustodelete:=l.head
for previoustodelete.next.data !=value{
	previoustodelete=previoustodelete.next
}
previoustodelete.next=previoustodelete.next.next
l.length --
}


func main() {

	mylist:=linkedlist{}
	node1:=&node(data:4)
	node2:=&node(data:5)

	mylist.prepend(node1)
	mylist.prepend(node2)

	mylist.printlistdata()
	
}