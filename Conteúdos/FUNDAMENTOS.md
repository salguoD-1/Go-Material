# Fundamentos da Linguagem Go

## Variáveis

- Variáveis são espaços na memória que armazenam valores. Para declarar uma variável, usamos a palavra reservada var seguida do nome da variável, o tipo e o valor. Além disso, podemos declarar variáveis utilizando inferência de tipos, ou seja, não precisamos declarar o tipo da variável, o compilador irá inferir o tipo da variável automaticamente.

```go
package main

import "fmt"

func main() {
	// Há duas formas de declarar uma variável:
	// 1. Declarando a variável com o tipo de dado que ela vai armazenar.
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// 2. Declarar uma variável sem usar var e sem especificar o tipo de dado que ela vai armazenar.
	// O compilador vai inferir o tipo de dado que a variável vai armazenar.
	variavel2 := "Variável 2"
	fmt.Println(variavel2)
}
```

Note que usamos a sintaxe variavel2 := "Variável 2" para declarar a variável. Isso é chamado de inferência de tipos. O compilador irá inferir o tipo da variável automaticamente.

## Declarando múltiplas variáveis

- Podemos declarar múltiplas variáveis de uma vez só. Para isso, basta separar os nomes das variáveis com vírgula.

```go
package main

import "fmt"

func main() {
	// Há duas formas de declarar uma variável:
	// 1. Declarando a variável com o tipo de dado que ela vai armazenar.
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// 2. Declarar uma variável sem usar var e sem especificar o tipo de dado que ela vai armazenar.
	// O compilador vai inferir o tipo de dado que a variável vai armazenar.
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

  // Declarando múltiplas variáveis
  var (
    variavel3 string = "Variável 3"
    variavel4 string = "Variável 4"
  )

  fmt.Println(variavel3, variavel4)

  // Declarando múltiplas variáveis com inferência de tipos
  variavel5, variavel6 := "Variável 5", "Variável 6"
  fmt.Println(variavel5, variavel6)
}
```

Note que declaramos múltiplas variáveis utilizando a sintaxe var (variavel3 string = "Variável 3", variavel4 string = "Variável 4"). Isso é chamado de bloco de declaração de variáveis. Além disso, note que declaramos múltiplas variáveis utilizando a sintaxe variavel5, variavel6 := "Variável 5", "Variável 6". Isso é chamado de declaração curta de variáveis.

## Constantes

- Constantes são valores que não podem ser alterados durante a execução do programa. Para declarar uma constante, usamos a palavra reservada const seguida do nome da constante, o tipo e o valor.

```go
package main

import "fmt"

func main() {
  // Declarando uma constante
  const PI float64 = 3.1415
  fmt.Println("O valor de PI é", PI)

  // Declarando uma constante sem especificar o tipo
  const PI2 = 3.1415
  fmt.Println("O valor de PI2 é", PI2)
}
```

## Tipos de dados

- Go é uma linguagem fortemente tipada e estática. Isso significa que todas as variáveis devem ter um tipo de dado e esse tipo de dado não pode ser alterado durante a execução do programa. Além disso, Go é uma linguagem compilada, ou seja, o código fonte é compilado para um código binário que pode ser executado em qualquer sistema operacional.

- Go possui os seguintes tipos de dados:

  - bool: representa um valor booleano (true ou false)
  - string: representa uma sequência de caracteres
  - int: representa um número inteiro
  - int8: representa um número inteiro com 8 bits
  - int16: representa um número inteiro com 16 bits
  - int32: representa um número inteiro com 32 bits
  - int64: representa um número inteiro com 64 bits
  - uint: representa um número inteiro sem sinal
  - uint8: representa um número inteiro sem sinal com 8 bits
  - uint16: representa um número inteiro sem sinal com 16 bits
  - uint32: representa um número inteiro sem sinal com 32 bits
  - uint64: representa um número inteiro sem sinal com 64 bits
  - byte: representa um byte
  - rune: representa um caractere
  - float32: representa um número de ponto flutuante com precisão simples
  - float64: representa um número de ponto flutuante com precisão dupla
  - complex64: representa um número complexo com precisão simples
  - complex128: representa um número complexo com precisão dupla

- Alguns exemplos com seus tipos de dados:

```go
package main

import "fmt"

func main() {
	// Para valores inteiros temos
	// int8, int16, int32, int64
	// Tamanho máximo de cada um deles
	var numero1 int8 = 127
	var numero2 int16 = 32767
	var numero3 int32 = 2147483647
	var numero4 int64 = 9223372036854775807
	fmt.Println(numero1, numero2, numero3, numero4)

	// Usando int ele vai utilizar o tamanho máximo que a arquitetura do seu computador suporta.
	var numeroInt int = 9223372036854775807;
	fmt.Println(numeroInt)

	// Usando uint para valores positivos
	// uint8, uint16, uint32, uint64
	// Tamanho máximo de cada um deles
	var numero5 uint8 = 255
	var numero6 uint16 = 65535
	var numero7 uint32 = 4294967295
	var numero8 uint64 = 18446744073709551615
	fmt.Println(numero5, numero6, numero7, numero8)


	// Para valores reais temos
	// float32, float64
	// A diferença entre eles é a precisão que eles tem. O float32 tem menos precisão que o float64.
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 1234567890.1234567890
	fmt.Println(numeroReal1, numeroReal2)

	// Para strings temos
	var str string = "Texto"
	fmt.Println(str)

	// Para booleanos temos
	var booleano1 bool = true
	var booleano2 bool = false
	fmt.Println(booleano1, booleano2)

	// Para caracteres temos
	// Em go não existe o tipo char, mas podemos usar o tipo rune(aliase) para representar um caractere. Usamos aspas simples para representar um caractere.
	var char1 byte = 'a'
	fmt.Println(char1)

	// Valores não inicializados
	// Quando não inicializamos uma variável, ela recebe o valor zero do tipo de dado que ela armazena.
	var numeroNaoIniciado int
	var strNaoIniciada string
	var booleanoNaoIniciado bool
	var charNaoIniciado byte
	// 0  false 0
	fmt.Println(numeroNaoIniciado, strNaoIniciada, booleanoNaoIniciado, charNaoIniciado)

	// Tipo de dado error

	// O tipo de dado error é um tipo de dado que representa um erro. Ele é uma interface, ou seja, ele é um tipo de dado que possui métodos.
	var erro error
	// Retorna nil.
	fmt.Println(erro)
}
```

[Voltar](../README.md)
