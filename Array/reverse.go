package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array is",arr)
	var arr2 [len(arr)]int
	for i:=0;i<len(arr);i++{
		arr2[(len(arr)-i-1)]=arr[i]
	}
	fmt.Println("Reversing is",arr2)
}
