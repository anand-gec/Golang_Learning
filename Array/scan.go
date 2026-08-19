package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's Your name :")
	//it is a single variable scanner also it skip the value that to outer to spacing
	// var name string;
	// fmt.Scan(&name)

	reader:=bufio.NewReader(os.Stdin)
	name,_:=reader.ReadString('\n')
	fmt.Println("Hello Mr.",name)


	
}