package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hare to call employee data")
	employee1Data()
	employee2Data()
	employee3Data()

}

type Employee struct {
	Name       string
	Age        int
	Salary     int
	Department string
}

func employee1Data() {
	var emp1 Employee
	emp1.Name = "Mohan"
	emp1.Age = 25
	emp1.Salary = 15000
	emp1.Department = "Fitter"
	fmt.Println("Employee1 data is ", emp1)

}

func employee2Data() {
	var emp2 Employee
	emp2.Name = "Sumit"
	emp2.Age = 26
	emp2.Salary = 18000
	emp2.Department = "Mechanic"
	fmt.Println("Employee2 data is ", emp2)
}

func employee3Data() {
	var emp3 Employee
	emp3.Name = "Aman"
	emp3.Age = 32
	emp3.Salary = 30000
	emp3.Department = "Team Leader"
	fmt.Println("Employee3 Data is ", emp3)
}
