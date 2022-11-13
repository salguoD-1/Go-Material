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



[Voltar](../README.md)
