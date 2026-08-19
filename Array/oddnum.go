package main

import "fmt"

func main() {
	arr := [...]int{4, 77, 123, 43, 87, 11, 21, 43, 56}
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			fmt.Printf("Odd number in array is %d", arr[i])
		}
	}
}
