package main

import (
	files "filess/connecting"
	"fmt"
)

type user struct {
	Name string
	Age  int
	Vill string
}

func main() {

	var emp user
	emp.Name = "Vikash"
	emp.Age = 23
	emp.Vill = "Basbariya"

	fmt.Println("Friend is ", emp)

	files.PrintMessage("really connecting file to print message")
	files.Add(2, 3)
	files.Multi(4, 2)

}
