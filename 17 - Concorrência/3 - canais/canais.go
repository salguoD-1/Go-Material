package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá, mundo!", canal)

	// O for percorre todas as execuções do canal e exibe o resultado na tela.
	for mensangem := range canal {
		fmt.Println(mensangem)
	}

	fmt.Println("Fim do programa!")
}


func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// "Canal recebe texto"
		canal <- texto
		time.Sleep(time.Second)
	}
	// Após a execução do for, fechamos o canal
	close(canal)
}