package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//File name is "Files"
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("While Creating file this error wii form.", err)
		return
	}
	defer file.Close()

	//to write in file
	content := "b. Each task doesn't wait for the other to finish before starting their work."
	_, errors := io.WriteString(file, content)
	if errors != nil {
		fmt.Println("Error to use function in to write the data", errors)
		return
	}
	fmt.Println("Successfully created the file and also write some lines.")
}
