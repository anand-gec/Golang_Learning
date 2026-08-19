package files

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)

	fmt.Println("Hello i an from the really connecting file to print message")

}

func Add(x, y int) int {

	sum := x + y
	fmt.Printf("addition of %d and %d is ", x, y)
	fmt.Println(sum)

	return sum
}

func Multi(x, y int) int {
	multi := x * y
	fmt.Printf("addition of %d and %d is ", x, y)
	fmt.Println(multi)
	return multi
}
