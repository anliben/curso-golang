package main

import "fmt"

func main() {
	fmt.Println("Hello, world! main")
	welcome()

	adicao, subtracao, multiplicacao := soma(1, 2)
	fmt.Println(adicao, subtracao, multiplicacao)
}

func soma(a int, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func welcome() {
	fmt.Println("Hello, world! utils")
}
