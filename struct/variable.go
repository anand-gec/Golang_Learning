package main

import "fmt"

type user struct {
	Name string
	Age  int
	Vill  string
}

func main() {

	var emp user
	emp.Name = "Vikash"
	emp.Age = 23
	emp.Vill = "Basbariya"

	fmt.Println("Friend is ", emp)

}
