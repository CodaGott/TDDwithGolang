package main

import "fmt"

func main() {
//count()
	duplicates := []string{"a", "a", "b", "c", "c", "d", "d", "e"}
	var noDuplicate []string
	j := 0

	for i := 0; i < len(duplicates) - 1; i++ {
		if duplicates[i] != duplicates[i + 1] {
			noDuplicate[j] = duplicates[i]
			fmt.Println("No duplicates", noDuplicate)
			j++
		}
	}
	noDuplicate[j] = duplicates[len(duplicates) - 1]
	fmt.Println("No duplicates", noDuplicate)
}

func count() {

	words := "This is the  chapter"
	countNum := 0

	for i := 0; i < len(words); i++ {
		if words[i] == ' ' {
			continue
		}
		countNum++
		}
	fmt.Println("Number of Words counted", countNum)

	err := validate("Dozie", "Male")

	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}
}

type inputError struct {
	message string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField

}

func validate(name, gender string) error {
	if name == "" || name == " " || gender == "" || gender == " " {
		return &inputError{"Name and Gender is mandatory", "name"}
	}else {
		fmt.Println("Provided name:", name, "and provided gender:",gender)
	}

	if gender == "" || gender == " " {
		return &inputError{"Gender is a mandatory field", "gender"}
	}

	return nil
}