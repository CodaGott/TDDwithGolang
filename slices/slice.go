package main

import (
	"fmt"
	"runtime"
)

func main() {
	slice := [4] string {"1st element,", "2nd element", "3rd element", "4th and last element"}

	fromSlice := slice[1:2]

	fmt.Println("the slice elements",slice)
	fmt.Println("the fromSlice elements",fromSlice)
	fmt.Println()
	fmt.Println("Capacity of the slice",cap(slice))
	fmt.Println("Capacity of the slice from",cap(fromSlice))
	fmt.Println("Length of the slice",len(slice))

	test :=[]int{7,3,5,5,2}
	fmt.Println("test",test)

	test2 := test[0:2]

	test3 := test2[0:]

	fmt.Println()

	test2[0] = 4
	fmt.Println("test2",test2)
	fmt.Println("test",test)
	fmt.Println("test3 ", test3)


	letters := [3]string {"a", "b", "c"}

	for i, letter := range letters {
		if letter == "c" {
			fmt.Println("the letter",letter, "is at position",i)
		}
	}

	for i := 0; i < len(letters); i++ {
		if letters[i] != "d" {
			fmt.Println("D not found in letters at position", letters[i])
		}
	}

	ebuka := student{"ebuka", "0900"}
	emma := student{"emma", "0900"}
	gold := student{"gold", "0900"}


	studentStuck := []student{ebuka, emma}

	fmt.Println("\nStudents before append", studentStuck)

	studentStuck = append(studentStuck, gold)

	fmt.Println("\nStudents after append", studentStuck)

	for i := 0; i < len(studentStuck); i++ {
		if studentStuck[i] == ebuka{
			println("\nfound")
		}
		fmt.Println(studentStuck[i])
	}


	src := []int{1,2,3,4,7, 9, 9, 8}
	dst := make([]int, 5)

	numOfElement := copy(dst, src)

	num :=[][]int{{3,7,8},{2,2,3}}

	fmt.Println("\n Number of copied elements", numOfElement, num)

	accountMap := map[string]account{}

	ebukaAccount := account{
		"002",
		"ebuka",
	}

	ebukaAccount2 := account{
		"002",
		"ebuka2",
	}

	accountMap["002"] = ebukaAccount



	accountMap["002"] = ebukaAccount2

	fmt.Println("Account map", accountMap)

	fmt.Println("Constant", a)
	fmt.Println("Constant", b)
	fmt.Println("Constant", c)
	fmt.Println("Constant", d)
	fmt.Println("Constant", e)

	fmt.Println(runtime.NumCPU())
}

type student struct {
	name string
	phoneNum string

}

const (
	a  = iota
	b
	c
	d
	e
)

type account struct {
	accountNum string
	accountOwner string
}