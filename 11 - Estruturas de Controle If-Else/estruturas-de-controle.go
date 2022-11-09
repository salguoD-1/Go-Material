package main

import "fmt"

func main() {
	// Declaramos uma variável idade do tipo int e atribuimos um valor a ela.
	idade := 18

	// Analisa se a condição é verdadeira, caso seja, executa o bloco if, caso contrário executa o bloco else.
	if idade >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade")
	}

	// Atribuindo um novo valor a variável idade
	idade = 8

	// Primeiro declaramos uma variável e inicializamos ela, nesse caso com a variável idade.
	// Em seguida usamos ponto e vírgula e fazemos a condição (outraIdade < 18)
	// Caso seja true, executa o bloco abaixo.
	if outraIdade := idade; outraIdade >= 18 {
		fmt.Println("Maior de idade!")
	} else if outraIdade < 10 {
		// Exibe: Você tem menos de 10 anos!
		fmt.Println("Você tem menos de 10 anos!")
	} else {
		fmt.Println("Menor de idade!")
	}
}