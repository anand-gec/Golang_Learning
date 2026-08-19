package main

import "fmt"

func main() {
	arr := [...]int{3, 6, 4, 9, 5, 8, 9, 7}
	largest := arr[0]
	second_largest := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] > second_largest && arr[i] != largest {
			second_largest = arr[i]
		}
	}
	fmt.Println("Second_largest value is ",second_largest)
}
