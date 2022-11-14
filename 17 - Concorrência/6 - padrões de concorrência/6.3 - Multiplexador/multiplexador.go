package main

import (
	"fmt"
	"time"
)

func main() {
	// Chamamos a função multiplexar passando a função escrever duas vezes, elas retornam um canal cada uma
	canal := multiplexar(escrever("Olá, mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}


// Passamos dois canais de saída do tipo string que retorna um canal do tipo string
func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <-chan string {
	// Canal de saída
	canalDeSaida := make(chan string)


	// Analisamos se os canais de entrada estão passando algum dado, caso estejam enviam para o canal de saída.
	go func() {
		for {
			select {
			case mensagem := <- canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	// Retorna o canal de saída
	return canalDeSaida
}

// Retorna um canal do tipo string
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}