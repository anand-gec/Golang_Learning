package main

import "fmt"

func main() {
	number := 20

	fmt.Println(number)
	fmt.Println(&number)
	case1(number)
	fmt.Println("number in case1",number)
	fmt.Println("number in case1",&number)
	case2(&number)
	fmt.Println("number in case2",number)
	fmt.Println("number in case2",&number)
}

func case1(number int){
	number=number+10
}

func case2 (number *int){
	*number=*number + 20
}
