package main

import "fmt"

type name2 struct {
		name4 int
		name3 *name2
	}

func main() {
	
for i:=0; i<5;i++{
	name3:=i
	fmt.Println("value is name3 ",name3)
}
for i:=7; i<9;i++{
	name4:=i
	fmt.Println("value of name4 ",name4)
}

// var name3 =5
// 	fmt.Println("name3 is ", name3)

node2:=3
fmt.Println("node 2 is ",node2)

}
