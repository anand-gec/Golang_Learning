package main

import "fmt"

type Vertex struct {
	X int
	Y int
	Z string
}

type User struct {
	Username string
	Age      int
	IsActive bool
}

func main() {

	var category Vertex
	category.Z = "Aman"
	category.X = 4

	fmt.Println("Name is ", category)
	fmt.Println(Vertex{1, 2, "Vikash"})

	// Declaring and initializing the struct type
	u := User{
		Username: "Alice",
		Age:      30,
		IsActive: true,
	}

	fmt.Println(u.Username) // Accessing fields using dot notation

	var emp1 User
	emp1.Username = "Mohan"
	emp1.Age = 22
	emp1.IsActive = true

	fmt.Println("data of Employee1 is", emp1)

var emp2 = User {
	Username: "Rohit",
	Age: 24,
	IsActive: false,
}

fmt.Println("Employee2 data is ",emp2)

}
