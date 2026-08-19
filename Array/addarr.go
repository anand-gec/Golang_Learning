// to add value in array
package main

import "fmt"

func main() {
	var arr [5]int
	var num int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Enter number to store in Array at position of arr[%d]:\n", i)
		fmt.Scan(&num)
		arr[i] = num
	}
	fmt.Println("arr is ", arr)
	for index, value := range arr {
		fmt.Printf("index of array is arr[%d] and value is %d\n", index, value)
	}
}
