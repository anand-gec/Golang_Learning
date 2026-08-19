package main

import "fmt"

func main() {
	fmt.Println("Enter number to i give you my friend name")
	var num int
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("Hey i am Vikash")
	case 2:
		fmt.Println("Hey i am Ranjesh")
	case 3:
		fmt.Println("Hey i am Ajay")
	case 4:
		fmt.Println("Hey i am Sangam")
	default:
		fmt.Println("By... ! This is out of my range..")
	}
}
