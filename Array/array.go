package main

import "fmt"

func main() {
	arr1 := [...]int{6, 8, 9, 2, 5, 7}
	fmt.Println("Array 1 is :", arr1)
	for i, j := 0, len(arr1)-1; i < j; i, j = i+1, j-1 {
		arr1[i], arr1[j] = arr1[j], arr1[i]
	}
	fmt.Println("After reverse the array ", arr1)
}
