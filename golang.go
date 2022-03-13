package main

import (
	"awesomeProject/methods"
	//"encoding/json"
	"fmt"
)

func main() {
	areaF := getAreaFunc()
	print(3, 4, areaF)

	fmt.Println("Adding numbers", add(3, 7, 9 , 10))


	emp1 := methods.Employee{
		Name:        "Dozie",
		Age:         "31",
		PhoneNumber: "08062611811",
		Salary:      44.00,
	}



	emp1.Details()
	fmt.Println()
	emp1.SetName("Chidozie")

	//



	//fmt.Println("Name now is",emp1.Name)

	emp1.Details()




}

type area func(int, int) int

func print(x, y int, a area) {
	fmt.Printf("Area is: %d\n", a(x,y))
}

func getAreaFunc() area {
	return func(x , y int) int {
		return x * y
	}
}

func add(numbers ...int) int {
	sum := 0

	for _, num := range numbers{
		sum += num
	}
	return sum
}

