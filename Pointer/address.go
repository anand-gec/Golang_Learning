package main

import "fmt"

func main() {
	// var name string = "vikash"
	// // var add *string
	// add := &name

	// fmt.Println("name is ", name)
	// fmt.Println("Address is ", add)
	// fmt.Println("* of Address is ", *add)
	// fmt.Println("& of Address is ", &add)

	// fmt.Println("& to print name Address is", &name)
	// *add = "Ajay"
	// fmt.Println("ajay is *add", name)
	// fmt.Println("ajay add", *add)
	// fmt.Println("ajay position add", &add)

	name := "RAVA"
	var add *string
	add =  &name
	fmt.Println("name is ", name)
	fmt.Println("add is ", &name)
	fmt.Println("val is ", add)
	fmt.Println("*val is ", *add)
	fmt.Println("&val is ", &add)
	fun2(&name)
	fmt.Println("name in function 2 is",name)
	fun1(&name)
	fmt.Println("name in function 1 is",name)
}
func fun1(name *string){
		*name = "Vikash"
}

func fun2(name *string){
	*name="Ajay"
}
