package main

import "fmt"

func main() {
	arr1 := [...]int{3, 56, 9, 3, 2, 7}
	arr2 := [...]int{6, 8, 3, 2, 2, 5, 0, 5}
	var arr3 [(len(arr1)+len(arr2))]int

	for i := 0; i <= len(arr1)-1; i++ {
		arr3[i] = arr1[i]
	}
	for i := 0; i <= len(arr2)-1; i++ {
		arr3[len(arr1)+i] = arr2[i]
	}
	// for i:=0;i<(len(arr1)+len(arr2));i++ {
	// 	fmt.Println("Final array3 is ", arr3[i])
	// }
	fmt.Println("final array merged is ",arr3)
	
}
