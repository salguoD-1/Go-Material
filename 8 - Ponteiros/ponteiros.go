package main

import "fmt"

func main() {
	var variavel1 int = 10
	// Armazena o valor da variável1
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	
	// Incrementamos uma unidade a variável1
	variavel1++
	// A variável2 não será 11 pois a variável1 é apenas uma cópia.
	fmt.Println(variavel1, variavel2)

	// Para resolver esse problema podemos utilizar ponteiros, que são uma referência de memória.
	var variavel3 int
	// Criamos um ponteiro utilizando a notação * na frente do tipo.
	var ponteiro *int

	variavel3 = 100
	// O operador & serve para indicar o lugar da memória em que a variável3 está armazenada.
	ponteiro = &variavel3
	// Exibe o endereço de memória da variável3
	fmt.Println(variavel3, ponteiro)

	// Para exibir o valor da variável, basta colocar um * na variável ponteiro
	fmt.Println(variavel3, *ponteiro) // Exibe 100
}