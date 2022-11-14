package main

import "fmt"

func main() {
	// Canal com buffer com capacidade máxima = 2
	canal := make(chan string, 2)

	// Podemos passar até dois dados para o canal
	canal <- "Olá, mundo!"
	canal <- "Olá, go!"

	// Passamos os dois dados do canal para as variáveis abaixo
	mensagem1 := <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}