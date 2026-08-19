package main

import "fmt"

func main() {

	//largest integer
	var large int = 0
	arr := [...]int{1, 3, 9, 8, 4, 6}
	large = arr[0]
	for i := 1; i < len(arr); i++ {
		if large < arr[i] {
			large = arr[i]
		}
	}
	fmt.Println("Large array value is :", large)
}
