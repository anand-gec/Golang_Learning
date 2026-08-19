// this is not array slice
package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{2, 4, 7, 9, 9, 3, 4, 6}

	fmt.Println("first the array is ", arr)
	arr = append(arr, 12, 45, 23, 87, 65)
	fmt.Println("After append the array is ", arr)

	//To delete the slice length
	letters := []string{"a", "b", "c", "d", "e"}


	letters = slices.Delete(letters, 2, 3) //hare 2 is deleted position and 3 is to continue
	

	fmt.Println(letters) 
}
