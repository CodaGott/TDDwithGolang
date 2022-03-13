package main

import "fmt"

const sayHello = "Hello, "
const sayHelloInSpanish = "Hola, "
const sayHelloInFrench = "Bonjour, "

func Hello(name, language string) string {
	if name == "" && language == ""{
		return sayHello + "World"
	}
	if language == "Spanish" {
		return sayHelloInSpanish + name
	}

	if language == "French" {

		return sayHelloInFrench + name
	}
	return sayHello + name
}

func main() {
	fmt.Println(Hello("Dozie", "Spanish"))
}


