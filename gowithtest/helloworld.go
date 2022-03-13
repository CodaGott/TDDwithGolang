package main

import "fmt"

const sayHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		return sayHello + "World"
	}
	return sayHello + name
}

func main() {
	fmt.Println(Hello("Dozie"))
}


