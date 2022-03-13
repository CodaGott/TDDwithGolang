package main

import "fmt"

const sayHello = "Hello, "

func Hello(name string) string {
	return sayHello + name
}

func main() {
	fmt.Println(Hello("Dozie"))
}


