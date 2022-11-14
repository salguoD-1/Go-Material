package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	// Função anônima
	go func () {
		for {
			// "Espera" 0,5 segundos
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// Função anônima
	go func () {
		for {
			// "Espera" 2 segundos
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
			// Caso o canal1 esteja pronto para receber dados, printa na tela
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1)
			// Caso o canal2 esteja pronto para receber dados, printa na tela
		case mensagemCanal2 := <- canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}