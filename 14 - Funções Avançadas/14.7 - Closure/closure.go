package main

import "fmt"

func closure() func() {
	// Variável texto da função closure.
	texto := "Dentro da função closure!"

	// A função closure acessa a variável externa a ela.
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	// Armazenamos o resultado da função e chamamos a funcaoNova.
	funcaoNova := closure()
	funcaoNova()
}