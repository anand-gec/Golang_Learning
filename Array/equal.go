package main

import "fmt"

func main() {
	fmt.Println("Find who is same to array")
	arr1 := [...]int{1, 3, 5}
	arr3 := [...]int{5, 3, 1}
	arr2 := [...]int{1, 3, 5}

	if arr1 == arr2 {
		fmt.Println("array1 and array2 similar")
	} else {
		fmt.Println("arr1 and arr2 are not similar")
	}
	if arr2 == arr3 {
		fmt.Println("arr2 and arr3 are similar")
	} else {
		fmt.Println("arr2 and arr3 are not similar")
	}
	if arr1 == arr3 {
		fmt.Println("arr1 and arr3 are similar")
	} else {
		fmt.Println("arr1 and arr3 are not similar")
	}
}
