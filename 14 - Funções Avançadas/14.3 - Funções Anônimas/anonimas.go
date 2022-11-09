package main

import "fmt"

func main() {
	// Função anônima que exibe Hello, world! na tela.
	func () {
		fmt.Println("Hello, world!")
	}() // É obrigatório o uso do parênteses no final.

	// Usando parâmetros em funções anônimas
	// Declaramos o parâmetro e o tipo e passamos o argumento dentro dos parênteses.
	func (user string) {
		fmt.Println("Olá", user)
	} ("Douglas")

	// Podemos criar uma função anônima que retorna algo e armazenar ela em uma variável.
	retorno := func (texto string) string {
		// A função Sprintf formata a saída, note que passamos %s que significa que queremos exibir o valor do texto que é do tipo string.
		return fmt.Sprintf("Retornando: %s", texto)
	}("Olá, eu sou um texto :)")
	// Exibimos o resultado do retorno.
	fmt.Println(retorno)
}