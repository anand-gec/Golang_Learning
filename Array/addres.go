package main

import "fmt"

func main() {
	number := 15
	case1(number)
	fmt.Println("address of &number is ", &number)
	fmt.Println("after case1 number is ", number)
	case2(&number)
	fmt.Println("address of number is ", &number)
	fmt.Println("after case2 number is ", number)

	employee("Vikash",23) // "Vikash" and 23 are the arguments
}

//pass by value
func case1(number int) {
	number = number + 2
	// fmt.Println("Case 1 number is ", number)
	// fmt.Println("Case 1 number address is ", &number)
}

//pass by reference
func case2(number *int) {
	*number = *number + 3
	// fmt.Println("Case2 in number is ", *number)
	// fmt.Println("Case2 in number address is ", &number)
}


func employee(name string, age int) { // 'name' and 'age' are parameters
	fmt.Printf("name is %s and age is %d",name,age)
}