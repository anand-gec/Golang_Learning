package main

import "fmt"

func main() {
	arr := [...]int{4, 12, 5, 68, 35, 33, 31, 21, 29}
	fmt.Println("array is ", arr)
	var num int = 0
	for i := 0; i < len(arr); i++ {
		num = 0
		for j := 2; j < arr[i]; j++ {
			if arr[i]%j == 0 {
				num = 1
			}
		}
		if num == 0 && arr[i] > 1 {
			fmt.Println("prime number is ", arr[i])
		}
	}
}
