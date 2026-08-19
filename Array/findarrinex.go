package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{1, 2, 3, 4, 5}
	var num int = 0
	var flag int = 0
	fmt.Println("Find when array is exist or note ")
	fmt.Scanf("%d", &num)
	for i := 0; i < len(arr); i++ {
		if num == arr[i] {
			flag = 1
			fmt.Printf("Number is %d is array position of arr[%d]\n", num, i)
			break
		}
	}
	if flag == 0 {
		fmt.Println("Item not found in array")
	}
}
