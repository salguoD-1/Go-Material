package main

import "fmt"

// Passamos um parâmetro seguido de três pontos e o tipo de dado do parâmetro + o tipo de dado de retorno.
func soma(numeros ...int) int {
	// Criamos um contador
	total := 0
	// Usamos o for para percorrer o slice que é gerado pela quantidade de parâmetros.
	for _, numero := range numeros {
		// Somamos o numero no parâmetro com o total.
		total += numero
	}
	// Retornamos o total.
	return total
}

// Função que recebe uma string e uma sequência de números.
func escrever(texto string, numeros ...int) {
	// Percorre o slice numeros e exibe na tela a string e o numero.
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// Chamamos a função passando 11 parâmetros e armazenamos o resultado.
	totalDaSoma := soma(1, 5, 10, 20, 100, 350, 3, 2, 1, 6, 123)

	// Exibe 621.
	fmt.Println(totalDaSoma)

	// Chamamos a função e passamos os dois argumentos.
	escrever("Olá, mundo!", 10, 20, 30, 40, 50)
}