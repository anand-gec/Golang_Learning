package main

import (
	"fmt"
)

func main() {
	arr2 := [...]int{1, 2, 3, 4}
	var arr5 [len(arr2)]int

	j := 0
	for i := 0; i <=len(arr2)-1; i++ {
		multi := arr2[i] * arr2[i]
		arr5[j] = multi
		j++
		// fmt.Println(multi)
		// }
	}
	fmt.Println("Array of squire is ", arr5)
}
