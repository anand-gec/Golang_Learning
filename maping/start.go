package main

import "fmt"

func main() {
	var temperature = make(map[string]int)
	//Create
	temperature["Gaya"] = 39
	temperature["Madhubani"] = 36
	temperature["Darbhanga"] = 37
	temperature["Patna"] = 38
	fmt.Println("Temperature is in Bihar", temperature)

	//update
	temperature["Gaya"] = 42
	fmt.Println("Gaya temperature after updated", temperature["Gaya"])

	//Delete
	delete(temperature, "Gaya")

	//Read
	fmt.Println("Temperature after delete Gaya is ", temperature)

	//check if exist or not
	_, check := temperature["Gaya"]
	fmt.Println("Is Gaya exist :", check)

	//not to be assigned
	// m:=map[string]interface{}
	// m["name"]="Vikash"

	// fmt.Println("nil map",m)

	for index, value := range temperature {
		fmt.Printf("In %q Temperature is %d\n", index, value)
	}

	var studentGrade = make(map[string]interface{})
	studentGrade["Vikash"] = 80
	studentGrade["Ajay"] = 85
	studentGrade["Ranjesh"] = 90

	fmt.Println("Student name and grade is \n", studentGrade)

	for index, value := range studentGrade {
		fmt.Printf("Student %q marks is :%d\n", index, value)
	}

	//check if exist or not 
	grade, exist := studentGrade["Prince"]
	fmt.Println("Grade of prince", grade)
	fmt.Println("is prince exist in table ", exist)

}
