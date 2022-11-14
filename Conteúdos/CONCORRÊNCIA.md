# Concorrência em Go

Concorrência é um conceito que diz respeito a execução de tarefas simultâneas. Em Go, a concorrência é implementada através de goroutines e canais.

Concorrência é diferente de paralelismo. Concorrência é a capacidade de executar várias tarefas simultaneamente. Paralelismo é a capacidade de executar várias tarefas simultaneamente em diferentes CPUs.

No exemplo abaixo nós temos uma função que é executada infinitamente, porém realizamos duas chamadas dessa função, no entanto, a segunda chamada nunca é executada pois ela aguarda a primeira chamada ser finalizada.

Mais sobre concoorrência em Go: [https://golang.org/doc/effective_go.html#concurrency](https://golang.org/doc/effective_go.html#concurrency)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	escrever("Olá, mundo!")
	escrever("Programando em Go!")
}


func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
```

Para resolver esse problema, podemos utilizar goroutines. Uma goroutine é uma função que é executada de forma independente da função principal. Para criar uma goroutine, basta utilizar a palavra reservada `go` antes da chamada da função.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá, mundo!")
	escrever("Programando em Go!")
}


func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
```

Agora, a segunda chamada da função `escrever` é executada de forma independente da primeira chamada. Isso significa dizer que a segunda chamada não aguarda a primeira chamada ser finalizada para poder ser executada.

## Usando waitgroups

Para aguardar a execução de goroutines, podemos utilizar o pacote `sync` e a estrutura `WaitGroup`. A estrutura `WaitGroup` é utilizada para aguardar a execução de goroutines.

```go
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
```

Note que usamos o pacote sync e a estrutura WaitGroup. A estrutura WaitGroup é utilizada para aguardar a execução de goroutines. Para isso, utilizamos o método `Add` para adicionar a quantidade de goroutines que irão rodar ao mesmo tempo. Após isso, utilizamos o método `Done` para remover uma unidade do "contador" `Add(n - 1)`. Por fim, utilizamos o método `Wait` para aguardar até que o "contador" esteja em 0. A cada execução da função escrever, o "contador" é decrementado em 1.

Ou seja, **As duas funções anônimas são executadas ao mesmo tempo, porém, o waitgroup aguarda a execução de ambas as funções `escrever()` para poder finalizar a execução do programa.** Isso resulta em uma sincronização entre as goroutines.

## Channels(Canais)

Channels são utilizados para comunicação entre goroutines(enviar/receber) dados. Para criar um channel, basta utilizar a palavra reservada `make` e passar o tipo do channel como parâmetro.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá, mundo!", canal)

	for {
		// A variável mensagem aguarda o canal receber um valor para poder armazenar os dados
		// Após fechar o canal, irá retorar 0 que significa false
		mensagem, aberto := <- canal

		// Confere se o canal está fechado, caso esteja sai do loop
		if !aberto {
			break
		}

		// Exibe os dados na tela
		fmt.Println(mensagem)
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
```

Note que criamos um channel do tipo string. Para enviar dados para o channel, utilizamos o operador `<-`. Para receber dados do channel, utilizamos o operador `<-` também. Para fechar o channel, utilizamos o método `close`. Para verificar se o channel está fechado, utilizamos a variável `aberto` que recebe o valor do channel. Caso o channel esteja fechado, a variável `aberto` recebe 0 que significa false. Caso o channel esteja aberto, a variável `aberto` recebe 1 que significa true.

**NOTA: Caso o canal não seja fechado, temos que a variável mensagem irá aguardar o canal receber um valor para poder armazenar os dados. Isso pode causar um `deadlock` que é um erro fatal `durante a execução do programa`.**

Para saber mais sobre deadlock, acesse os links: <br> [https://pt.wikipedia.org/wiki/Deadlock](https://pt.wikipedia.org/wiki/Deadlock) <br>
[https://programming.guide/go/detect-deadlock.html](https://programming.guide/go/detect-deadlock.html)

Uma outra forma de usar o for, seria fazendo uso do range.

```go
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
```

Com o range, não precisamos verificar se o canal está fechado, pois o range já faz isso por nós.

## Canais com Buffer

Podemos criar um channel com buffer, ou seja, podemos definir a quantidade de dados que o channel pode armazenar. Para isso, basta passar a quantidade de dados que o channel pode armazenar como parâmetro. Isso evita o deadlock.

```go
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
```

## Select

O select é utilizado para receber dados de vários channels. Para isso, utilizamos o método `default` para verificar se o channel está aberto. Caso o channel esteja aberto, o select irá receber os dados do channel. Caso o channel esteja fechado, o select irá ignorar o channel.

No exemplo abaixo nós temos um problema, note que o primeiro canal leva cerca de 0.5 segundos para ser executado, enquanto isso o segundo canal leva cerca de 2 segundos para ser executado, isso gera um **delay** no nosso último for, ou seja, o for irá aguardar o canal 2 ser executado para poder exibir os dados na tela. Para resolver esse problema, utilizamos o select.

```go
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

	// Esse loop infinito só pula para a próxima iteração após o último canal enviar seus dados(a cada 2seg)
	for {
		mensagemCanal1 := <- canal1
		fmt.Println(mensagemCanal1)

		mensagemCanal2 := <- canal2
		fmt.Println(mensagemCanal2)
	}
}
```

Usando a estrutura select:

```go
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
```

No exemplo acima, temos que a estrutura select é mais performática, pois ela não espera o canal 2 ser executado para poder exibir os dados na tela. Ou seja, o canal1 irá ser executado quatro vezes enquanto o canal2 é executado apenas uma vez nesse intervalo de dois segundos. A estrutura select é utilizada somente para concorrência.

## Padrões de Concorrência

Padrões de concorrência são formas de resolver problemas de concorrência.

### Padrão Worker Pool

O padrão worker pool é utilizado para resolver problemas de concorrência. Nesse padrão, temos um canal de entrada e um canal de saída. O canal de entrada recebe os dados que serão processados, enquanto o canal de saída recebe os dados processados. O canal de entrada é utilizado para enviar os dados para os workers, enquanto o canal de saída é utilizado para receber os dados processados dos workers.

```go
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
```

O padrão worker pool é muito utilizado para resolver problemas relacionados a filas. No exemplo acima nós temos uma fila de tarefas que são enviadas para os workers, que são os processos que executam a função fibonacci. O resultado dessas tarefas são armazenados no canal resultados e exibidos na tela.

Note que a função worker possui o canal tarefas que especificamos(usando <- antes do chan) como sendo um canal de entrada, ou seja, ele só recebe dados. Já o canal resultados é um canal de saída(especificamos usando <- depois do chan), ou seja, ele só envia dados. Isso é muito importante para evitar que os dados sejam enviados para o canal errado.

[Voltar](../README.md)
