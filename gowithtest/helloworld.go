package main

import "fmt"

const sayHello = "Hello, "
const sayHelloInSpanish = "Hola, "
const sayHelloInFrench = "Bonjour, "

func Hello(name, language string) string {
	if name == "" && language == ""{
		return sayHello + "World"
	}
	//if language == "Spanish" {
	//	return sayHelloInSpanish + name
	//}
	//
	//if language == "French" {
	//
	//	return sayHelloInFrench + name
	//}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix =  sayHelloInFrench
	case "Spanish":
		prefix = sayHelloInSpanish
	default:
		prefix = sayHello
	}
	return
}

func main() {
	fmt.Println(Hello("Dozie", "Spanish"))
}


