package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Pacote sync
	var waitGroup sync.WaitGroup

	// Adicionamos a quantidade de goroutines que irão rodar ao mesmo tempo.
	waitGroup.Add(2)

	// Função anônima que exibe algo na tela
	go func () {
		escrever("Olá, mundo!")
		// Método .Done() remove uma unidade do "contador" Add(n - 1).
		waitGroup.Done()
	}()
	
	// Função anônima que exibe algo na tela
	go func () {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	// Aguarda até que o "contador" esteja em 0.
	waitGroup.Wait()
}


func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}