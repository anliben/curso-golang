package package001

import "fmt"

const nome = "Gopher"

func Welcome() {
	fmt.Println("Initializing the application... ")
}

func main() {

	Welcome()

	// Examples of primitive data types
	var number int = 10
	var pi float64 = 3.14
	var isTrue bool = true
	var character byte = 'A'
	var text string = "Hello, World!"

	fmt.Println(number)
	fmt.Println(pi)
	fmt.Println(isTrue)
	fmt.Println(character)
	fmt.Println(text)
	fmt.Println("meu nome e ", nome)
}
