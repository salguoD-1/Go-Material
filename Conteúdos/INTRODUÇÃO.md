# Introdução a linguagem Go

# Por que aprender Go?

1. Go é uma linguagem extremamente eficiente
2. É uma linguagem compilada
3. É uma linguagem otimizada para usar mais de um núcleo do processador
4. É uma linguagem feita para liadr muito bem com concorrência
5. É uma linguagem com uma comunidade muito ativa
6. É uma linguagem com uma [documentação](https://go.dev/doc/) ótima
7. É uma linguagem com uma curva de aprendizado muito baixa
8. É uma linguagem com uma simplicidade incrível como Python ou Javascript, além disso, tem a robustez de uma linguagem como C# ou Java.
9. É uma linguagem com uma sintaxe muito próxima de C, C++ e Java, facilitando muito a transição de uma linguagem para outra.
10. É uma linguagem criada pela Google.

## Hello World em Go

```go
// Pacote main
package main

// Importa o pacote fmt
import (
	"fmt"
)
// Função principal
func main() {
	fmt.Println("Hello, World!")
}
```

## O que são pacotes?

Pacotes são um conjunto de arquivos que contém funções, variáveis, constantes, tipos e interfaces que podem ser usados por outros pacotes. O pacote `main` é o pacote principal de um programa Go. Ele contém a função `main()` que é a função principal de um programa Go. O pacote `fmt` é um pacote padrão do Go que contém funções para formatar e imprimir dados. Pense em pacotes como é o caso do package.json do Node.js.

## O que é import?

A palavra reservada import é usada para importar pacotes.

## O que é a função main?

A função main é a função principal de um programa em Go. Ela é a primeira função a ser executada quando o programa é iniciado.

## O que é fmt?

fmt é um pacote que contém funções para formatar e imprimir dados.

## O que é println?

println é uma função que imprime uma linha de texto na tela. Ela faz parte do pacote fmt.

## Criando um módulo para armazenar os nossos pacotes

```go
// mod significa a extenção do arquivo. O nome do módulo é modulo.
go mod init modulo
```

## Rodando o programa

- Podemos rodar o código diretamente usando o comando go run

```go
go run main.go
```

## Compilando o programa

- Ou podemos compilar o código usando o comando go build

```go
go build main.go
```

- Executando o programa compilado

```go
./main
```

## Criando um pacote auxiliar

- Criamos um pacote chamado auxiliar que contém um arquivo chamado auxiliar.go

```go
package auxiliar

// Permite usar o pacote fmt para escrever na tela
import "fmt"

// Função que escreve na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
}
```

- **NOTA: Em go devemos utilizar a primeira letra maiúscula para tornar a função pública, ou seja, acessível por outros pacotes. Caso contrário, a função será privada e só poderá ser acessada dentro do mesmo pacote. Além disso, é recomendado que tenha um comentário em cima de cada função descrevendo o que a função faz.**

- Importamos o pacote auxiliar no nosso arquivo main.go

```go
package main

// Importa o pacote fmt e o pacote auxiliar
import (
	"modulo/auxiliar"
	"fmt"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()
}
```

[Voltar](../README.md)
