package methods

import "fmt"

type Employee struct {
	Name string
	Age string
	PhoneNumber string
	Salary float64
}

func (e *Employee) SetName(newName string) {
	e.Name = newName
}

func (e *Employee) Details() {
	fmt.Println("Name of employee", e.Name)
	fmt.Println("Age of employee", e.Age)
	fmt.Println("Phone number of employee", e.PhoneNumber)
	fmt.Println("Salary of employee", e.Salary)
}


