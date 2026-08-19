package main

import "fmt"

func main() {
	var name string= "Vikash"
	fmt.Println("name is nothing changed",name)  //name is nothing changed Vikash

	ptr := &name
	sec_ptr := &ptr

	*ptr = "Ranjesh"
	fmt.Println("name is ", name)  //name is  Ranjesh
	fmt.Println("name is &ptr", *ptr)  //name is &ptr Ranjesh
	fmt.Println("name is &ptr", &ptr)  //name is &ptr 0x35cae6dea060
	
	**sec_ptr = "Ajay"
	fmt.Println("name is ", name)  //name is  Ajay
	fmt.Println("name is &ptr", *ptr)  //name is &ptr Ajay
	fmt.Println("name is &ptr", &ptr) //name is &ptr 0x35cae6dea060
	fmt.Println("name is &ptr", **sec_ptr) //name is &ptr Ajay
	fmt.Println("name is &ptr", &sec_ptr) //name is &ptr 0x35cae6dea068
}
