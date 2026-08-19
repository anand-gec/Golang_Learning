package main

import "fmt"

func main() {
	fmt.Println("Hare to call functions")
	ans, err := div(4, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)
	fmt.Println("Hare to 2nd call functions")
	multti, _ := multi(4, 0)
	fmt.Println(multti)

}

func div(x, y float64) (float64, error) {
	if x == 0 || y == 0 {
		return 0, fmt.Errorf("You are putting zero in element")
	}
	return x / y, nil
}

func multi(x, y float64) (float64, string) {
	if x == 0 || y == 0 {
		return 0, "You are putting zero in element"
	}
	return x / y, "nil"
}
