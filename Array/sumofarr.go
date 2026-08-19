package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	// var summ int = 0
	// var arr3=[...]int{}
	arr := [...]int{2, 7, 4, 9, 3}
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("1st Array is:", arr)
	fmt.Println("2nd Array is:", arr2)

	for i := 0; i <= len(arr)-1; i++ {
		sum = sum + arr[i]
	}
	fmt.Println("Sum of all array value is ", sum)
}
