package main

import "fmt"

func main() {
	arr := [...]int{2, 8, 3, 6, 5}
	for index, value := range arr {
		fmt.Printf("index of number is %d and value is %d\n", index, value)
	}
}
