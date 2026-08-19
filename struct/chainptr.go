package main

import "fmt"

func main() {
	var num int = 10
	// var ptr *int
	// var ptrr **int

	ptr := &num
	ptrr := &ptr
	fmt.Println("normal case", num)
	fmt.Println("*ptr1", *ptr)
	fmt.Println("**ptr2", **ptrr)

	**ptrr = 20
	fmt.Println("normal case", num)
	fmt.Println("*ptr1", *ptr)
	fmt.Println("**ptr2", **ptrr)

	*ptr = 30
	fmt.Println("normal case", num)
	fmt.Println("*ptr1", *ptr)
	fmt.Println("**ptr2", **ptrr)

	num = 50
	fmt.Println("normal case", num)
	fmt.Println("*ptr1", *ptr)
	fmt.Println("**ptr2", **ptrr)
}
