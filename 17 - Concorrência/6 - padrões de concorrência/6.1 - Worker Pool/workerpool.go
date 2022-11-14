package main

import "fmt"

func main() {
	// Criamos dois canais com 45 posições
	// Armazena as tarefas
	tarefas := make(chan int, 45)
	// Armazena os resultados
	resultados := make(chan int, 45)
	
	// Usando quatro processos para executar a nossa função, isso reduz o tempo de execução.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Passamos o valor de i para o canal tarefas
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	// Fechamos o canal tarefas para evitar receber mais dados.
	close(tarefas)

	// Armazenamos o resultado do canal resultados e exibimos na tela
	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}

}

// Especificando que o canal tarefas só recebe dados enquanto que o canal resultados só envia dados.
func worker(tarefas <-chan int, resultados chan<- int) {
	// Para cada tarefa nós chamamos a função fibonacci passando a tarefa e armazenando o resultado no canal resultados.
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

// Função fibonacci recursiva
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}