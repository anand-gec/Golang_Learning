package main

import (
	"fmt"
	"slices"
)

func main() {
	// arr:= make([]int, 3, 6)
	arr := []int{2, 4, 7, 9, 9, 3, 4, 6}
	arr2 := make([]int, 2, 6) // 2 is to start fill the data hare..
	arr2 = append(arr2, 3)
	arr2 = append(arr2, 37)
	arr2 = append(arr2, 7)
	arr2 = append(arr2, 1)
	arr2 = append(arr2, 9)
	arr2 = append(arr2, 10)
	arr2 = append(arr2, 8)

	fmt.Println("first the array is ", arr)

	arr = append(arr, 12, 45, 23, 87, 65)
	fmt.Println("After append the array is ", arr)

	//To delete the slice length
	letters := []string{"a", "b", "c", "d", "e"}

	letters = slices.Delete(letters, 2, 3) //hare 2 is deleted position and 3 is to continue
	// hare 2 is like to delete start where 3

	fmt.Println(letters)
	fmt.Println("Second array is ",arr2)
}
