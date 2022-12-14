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

## Funções

- Funções são blocos de código que realizam uma tarefa específica. Para declarar uma função, usamos a palavra reservada func seguida do nome da função, os parâmetros e o tipo de dado que a função retorna.

```go
package main

import (
	"fmt"
)

// Palavra-chave func, nome da função, (n-parâmetros com cada tipo) retorno da função.
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função principal
func main() {
	// Armazena o resultado da função.
	soma := somar(10, 20);
	// Exibe 30.
	fmt.Println(soma)
}
```

Note que a função somar recebe dois parâmetros do tipo int8 e retorna um valor do tipo int8. Para declarar uma função que não retorna nada, usamos o tipo de dado void.

## Armazenando uma função em uma variável

Funções podem ser armazenadas em variáveis. Para isso, basta **declarar uma variável do `tipo de dado func` e atribuir uma função a ela.**

```go
package main

import (
	"fmt"
)

// Palavra-chave func, nome da função, (n-parâmetros com cada tipo) retorno da função.
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função principal
func main() {
	// Armazena o resultado da função.
	soma := somar(10, 20);
	// Exibe 30.
	fmt.Println(soma)

	// Armazenando uma função em uma variável
	var funcao = func() {
		fmt.Println("Eu sou uma função!")
	}

	// Chamando a função
	funcao()
}
```

- Usando parâmetros e retornando um resultado

```go
package main

import (
	"fmt"
)

// Palavra-chave func, nome da função, (n-parâmetros com cada tipo) retorno da função.
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função principal
func main() {
	// Armazena o resultado da função.
	soma := somar(10, 20);
	// Exibe 30.
	fmt.Println(soma)

	// Armazenando uma função em uma variável
	var funcao = func() {
		fmt.Println("Eu sou uma função!")
	}

	// Chamando a função
	funcao()

	// A função funcao2 tem como parâmetro uma string e retorna uma string.
	var funcao2 = func(txt string) string {
		return txt;
	}

	// Chamamos a função, passamos uma string como argumento e armazenamos o seu resultado.
	result := funcao2("Oi, eu sou uma função com parâmetro e retorno de dado :)")
	// Exibe: Oi, eu sou uma função com parâmetro e retorno de dado :)
	fmt.Println(result)
}
```

## Funções em Go podem ter mais de um retorno

- Para declarar uma função que retorna mais de um valor, basta separar os tipos de dados que a função retorna com vírgula.

```go
package main

import (
	"fmt"
)

// Palavra-chave func, nome da função, (n-parâmetros com cada tipo) retorno da função.
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Podemos declarar o tipo de dado do parâmetro no final.
// Passamos a quantidade e tipo de retorno da função entre parênteses.
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	// Retornamos os dois resultados.
	return soma, subtracao
}

// Função principal
func main() {
	// Armazena o resultado da função.
	soma := somar(10, 20);
	// Exibe 30.
	fmt.Println(soma)

	// Armazenando uma função em uma variável
	var funcao = func() {
		fmt.Println("Eu sou uma função!")
	}

	// Chamando a função
	funcao()

	// A função funcao2 tem como parâmetro uma string e retorna uma string.
	var funcao2 = func(txt string) string {
		return txt;
	}

	// Chamamos a função, passamos uma string como argumento e armazenamos o seu resultado.
	result := funcao2("Oi, eu sou uma função com parâmetro e retorno de dado :)")
	// Exibe: Oi, eu sou uma função com parâmetro e retorno de dado :)
	fmt.Println(result)

	// Para usar a função com retorno duplo, nós separamos as variáveis
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	// Exibe: 25 -5
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
```

- **NOTA:** Se não quisermos armazenar algum resultado tipo o resultadoSoma ou resultadoSubtracao, podemos usar o underline \_ para ignorar o resultado.

```go
package main

import (
	"fmt"
)

// Palavra-chave func, nome da função, (n-parâmetros com cada tipo) retorno da função.
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Podemos declarar o tipo de dado do parâmetro no final.
// Passamos a quantidade e tipo de retorno da função entre parênteses.
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	// Retornamos os dois resultados.
	return soma, subtracao
}

// Função principal
func main() {
	// Armazena o resultado da função.
	soma := somar(10, 20);
	// Exibe 30.
	fmt.Println(soma)

	// Armazenando uma função em uma variável
	var funcao = func() {
		fmt.Println("Eu sou uma função!")
	}

	// Chamando a função
	funcao()

	// A função funcao2 tem como parâmetro uma string e retorna uma string.
	var funcao2 = func(txt string) string {
		return txt;
	}

	// Chamamos a função, passamos uma string como argumento e armazenamos o seu resultado.
	result := funcao2("Oi, eu sou uma função com parâmetro e retorno de dado :)")
	// Exibe: Oi, eu sou uma função com parâmetro e retorno de dado :)
	fmt.Println(result)

	// Para usar a função com retorno duplo, nós separamos as variáveis
	resultadoSoma, _ := calculosMatematicos(10, 15)
	// Exibe: 25
	fmt.Println(resultadoSoma)
}
```

Ou seja, para não armazenar o resultado da subtração, nós usamos o \_.

## Operadores em Go

- Go possui os operadores aritméticos, relacionais, lógicos e de atribuição.

### Operadores aritméticos

- São operadores que realizam operações matemáticas.

| Operador | Descrição     |
| -------- | ------------- |
| +        | Adição        |
| -        | Subtração     |
| \*       | Multiplicação |
| /        | Divisão       |
| %        | Módulo        |

- Exemplos

```go
package main

import (
	"fmt"
)

func main() {
	soma := 3 + 5
	subtracao := 10 - 3
	multiplicacao := 3 * 5
	divisao := 10 / 2
	modulo := 20 % 5

	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtracao)
	fmt.Println("Multiplicação: ", multiplicacao)
	fmt.Println("Divisão: ", divisao)
	fmt.Println("Módulo: ", modulo);
}
```

**NOTA:** Em go, não podemos fazer operações com tipos diferentes, por exemplo, não podemos fazer 10 + 5.5, pois 10 é um inteiro e 5.5 é um float64. Não podemos somar dois inteiros com tipos diferentes, ex: int16 e int32.

## Operador de atribuição

- São operadores que atribuem valores a variáveis.

| Operador | Descrição                    |
| -------- | ---------------------------- |
| =        | Atribuição                   |
| +=       | Atribuição com adição        |
| -=       | Atribuição com subtração     |
| \*=      | Atribuição com multiplicação |
| /=       | Atribuição com divisão       |
| %=       | Atribuição com módulo        |

- Exemplos

```go
package main

import (
  "fmt"
)

func main() {
  var numero int8 = 10
  numero += 20 // numero = numero + 20
  numero -= 10 // numero = numero - 10
  numero *= 2 // numero = numero * 2
  numero /= 4 // numero = numero / 4
  numero %= 3 // numero = numero % 3

  fmt.Println(numero)
}
```

## Operadores relacionais

- São operadores que realizam comparações entre valores.

| Operador | Descrição        |
| -------- | ---------------- |
| ==       | Igual a          |
| !=       | Diferente de     |
| <        | Menor que        |
| >        | Maior que        |
| >=       | Maior ou igual a |
| <=       | Menor ou igual a |

- Exemplos

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println(1 == 1) // true
  fmt.Println(1 != 1) // false
  fmt.Println(1 < 1) // false
  fmt.Println(1 > 1) // false
  fmt.Println(1 >= 1) // true
  fmt.Println(1 <= 1) // true
}
```

## Operadores lógicos

- São operadores que realizam comparações entre valores booleanos.

| Operador | Descrição |
| -------- | --------- |
| &&       | E         |
| \|\|     | OU        |
| !        | Negação   |

- Exemplos

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println(true && true) // true
  fmt.Println(true && false) // false
  fmt.Println(true || true) // true
  fmt.Println(true || false) // true
  fmt.Println(!true) // false
}
```

## Opeardores unários

- São operadores que realizam operações com apenas um operando.

| Operador | Descrição      |
| -------- | -------------- |
| +        | Sinal positivo |
| -        | Sinal negativo |
| ++       | Incremento     |
| --       | Decremento     |

- Exemplos

```go
package main

import (
  "fmt"
)

func main() {
  numero := 10
  fmt.Println(+numero) // 10
  fmt.Println(-numero) // -10
  numero++
  fmt.Println(numero) // 11
  numero--
  fmt.Println(numero) // 10
}
```

## Operador ternário

- O operador ternário é um operador que recebe três operandos, sendo o primeiro um valor booleano, o segundo um valor que será retornado caso o primeiro seja verdadeiro e o terceiro um valor que será retornado caso o primeiro seja falso.

- Exemplo

```go
package main

import (
  "fmt"
)

func main() {
  numero := 10
  var resultado string

  if numero > 5 {
    resultado = "Maior que 5"
  } else {
    resultado = "Menor que 5"
  }

  fmt.Println(resultado)

  // Usando o operador ternário
  resultado = numero > 5 ? "Maior que 5" : "Menor que 5"
  fmt.Println(resultado)
}
```

## Structs

- Structs são estruturas de dados que podem armazenar vários tipos de dados.

- Structs são o mais próximo que temos de classes em go.

- Exemplo

```go
package main

import (
	"fmt"
)

// Estrutura do tipo usuario
type usuario struct {
  nome string
  idade uint8
  email string
}

// Função principal
func main() {
	// Declaramos uma variável chamada user do tipo usuario.
  var user usuario
	// Usamos a notação dot(.) para acessar cada variável da estrutura e atribuimos um valor.
  user.nome = "Douglas"
  user.idade = 23
  user.email = "douglas@email.com"

	// Exibe o resultado de cada variável
	fmt.Println(user.email)
	fmt.Println(user.idade)
	fmt.Println(user.email)
}
```

No exemplo acima declaramos uma variável do tipo usuario, que é uma estrutura que contém três variáveis, nome, idade e email. Para acessar cada uma dessas variáveis, usamos a notação dot(.) e atribuímos um valor a cada uma delas. Para exibir o valor de cada variável, usamos a notação dot(.) novamente.

- **NOTA:** Caso não passemos nenhum valor para as variáveis da estrutura, temos que o resultado de cada variável será o valor padrão do tipo de dado que ela armazena.

- Exemplo

```go
package main

import (
  "fmt"
)

// Estrutura do tipo usuario
type usuario struct {
  nome string
  idade uint8
  email string
}

// Função principal
func main() {
  // Declaramos uma variável chamada user do tipo usuario.
  var user usuario

  // Exibe o resultado de cada variável
  fmt.Println(user.nome) // ""
  fmt.Println(user.idade) // 0
  fmt.Println(user.email) // ""
}
```

- Podemos declarar uma variável do tipo usuario e atribuir valores a cada uma das variáveis da estrutura usando a notação de inferência de tipo.

```go
package main

import (
  "fmt"
)

// Estrutura do tipo usuario
type usuario struct {
  nome string
  idade uint8
  email string
}

// Função principal
func main() {
  // Declaramos uma variável chamada user do tipo usuario e atribuímos valores a cada uma das variáveis da estrutura.
  user := usuario{
    nome: "Douglas",
    idade: 23,
    email: "douglas@email.com"
  }

  // Exibe o resultado de cada variável
  fmt.Println(user.nome) // Douglas
  fmt.Println(user.idade) // 23
  fmt.Println(user.email) // douglas@email.com
}
```

Note que a notação é semelhante a de um **objeto em JavaScript.**

## Podemos ter structs dentro de structs

- Exemplo

```go
package main

import (
  "fmt"
)

// Estrutura do tipo usuario
type usuario struct {
  nome string
  idade uint8
  email string
  endereco endereco
}

// Estrutura do tipo endereco
type endereco struct {
  logradouro string
  numero uint8
}

// Função principal
func main() {
  // Declaramos uma variável chamada user do tipo usuario e atribuímos valores a cada uma das variáveis da estrutura.
  user := usuario{
    nome: "Douglas",
    idade: 23,
    email: "douglas@email.com",
    // Atribuímos valores a cada uma das variáveis da estrutura endereco.
    endereco: endereco{
      logradouro: "Rua aprendendo Go",
      numero: 13,
    },
  }

  // Exibe o resultado de cada variável
  fmt.Println(user.nome) // Douglas
  fmt.Println(user.idade) // 23
  fmt.Println(user.email) // douglas@email.com
  fmt.Println(user.endereco.logradouro) // Rua aprendendo Go
  fmt.Println(user.endereco.numero) // 13
}
```

- Exibindo o resultado:

```bash
go run structs.go
Douglas
23
douglas@email.com
Rua aprendendo Go
13
```

- Note que passamos o valor de cada variável da estrutura endereco dentro da estrutura usuario. Para acessar cada uma das variáveis da estrutura endereco, usamos a notação dot(.) novamente. Note também que criamos uma variável chamada endereco do tipo endereco e atribuímos valores a cada uma das variáveis da estrutura.

## Heranças

- Go não tem herança, mas podemos simular heranças usando structs.

- Exemplo

```go
package main

import "fmt"

// Struct do tipo pessoa
type pessoa struct {
	nome string
	sobrenome string
	idade int8
	altura uint8
	peso float32
}

// Struct do tipo estudante
type estudante struct {
	// Passamos o nome da nossa struct que queremos "herdar". Não é necessário passar o tipo.
	pessoa
	curso string
	faculdade string
}

func main() {
	// Declaramos uma variável chamada pessoaUm do tipo pessoa e setamos alguns valores.
	pessoaUm := pessoa {
		"Douglas",
		"Cunha de Jesus",
		23,
		172,
		68.5,
	}

	// Exibe: {Douglas Cunha de Jesus 23 172 68.5}
	fmt.Println(pessoaUm)

	// Declaramos uma variável chamada estudanteUm do tipo estudante e passamos uma struct(pessoaUm) e setamos dois valores.
	estudanteUm := estudante {
		pessoaUm,
		"Sistemas de Informação",
		"Universidade Federal de Sergipe",
	}

	// Exibe: {{Douglas Cunha de Jesus 23 172 68.5} Sistemas de Informação Universidade Federal de Sergipe}
	fmt.Println(estudanteUm)

	// Para acessar as variáveis da struct pessoa basta usar a notação dot(.) seguido do nome da variável.
	fmt.Println(estudanteUm.nome)
	fmt.Println(estudanteUm.sobrenome)
	fmt.Println(estudanteUm.idade)
	fmt.Println(estudanteUm.altura)
	fmt.Println(estudanteUm.peso)
	fmt.Println(estudanteUm.curso)
	fmt.Println(estudanteUm.faculdade)
}
```

- O resultado será:

```bash
go run heranca.go
{Douglas Cunha de Jesus 23 172 68.5}
{{Douglas Cunha de Jesus 23 172 68.5} Sistemas de Informação Universidade Federal de Sergipe}
Douglas
Cunha de Jesus
23
172
68.5
Sistemas de Informação
Universidade Federal de Sergipe
```

- Ou seja, para usarmos "herança" em Go, basta passar uma struct como valor de uma variável de outra struct. Como é o caso da struct estudante que recebeu a struct pessoa como valor. Para exibir os valores das variáveis da struct pessoa, basta usar a notação dot(.) seguido do nome da variável.

## Ponteiros

- Ponteiros são variáveis que armazenam o endereço de memória de outra variável.

- Exemplo

```go
package main

import "fmt"

func main() {
	var variavel1 int = 10
	// Armazena o valor da variável1
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// Incrementamos uma unidade a variável1
	variavel1++
	// A variável2 não será 11 pois a variável1 é apenas uma cópia.
	fmt.Println(variavel1, variavel2)

	// Para resolver esse problema podemos utilizar ponteiros, que são uma referência de memória.
	var variavel3 int
	// Criamos um ponteiro utilizando a notação * na frente do tipo.
	var ponteiro *int

	variavel3 = 100
	// O operador & serve para indicar o lugar da memória em que a variável3 está armazenada.
	ponteiro = &variavel3
	// Exibe o endereço de memória da variável3
	fmt.Println(variavel3, ponteiro)

	// Para exibir o valor da variável, basta colocar um * na variável ponteiro
	fmt.Println(variavel3, *ponteiro) // Exibe 100
}
```

- Ou seja, ponteiros são úteis para quando queremos alterar o valor de uma variável que está em outro lugar da memória. Para isso é só criar um ponteiro usando a notação \* na frente do tipo da variável e atribuir o endereço de memória da variável que queremos alterar usando a notação \&. Para exibir o valor da variável, basta colocar um \* na variável ponteiro.

## Arrays e Slices

- Arrays e Slices são estruturas de dados que armazenam uma coleção de valores do mesmo tipo. Além disso, é necessário definir o tamanho do array ou slice.

- Exemplo

```go
package main

import "fmt"

func main() {
	// var nome_do_array[tamanho do array] tipo do array
	var array1[5] int
	// Exibe: [0 0 0 0 0]
	fmt.Println(array1)

	// Populando o array
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	// Exibe: [1 2 3 4 5]
	fmt.Println(array1)

	// Podemos declarar arrays usando inferência de tipo
	array2 := [5] string {
		"Posição1", "Posição 2", "Posição 3", "Posição 4", "Posição 5",
	}

	// Exibe: [Posição1 Posição 2 Posição 3 Posição 4 Posição 5]
	fmt.Println(array2)

	// Outra forma de declarar um array é usando o "operador spread" semelhante a do JavaScript que calcula automaticamente o tamanho do array.
	array3 := [...] int {
		1,
		2,
		3,
		4,
		5,
	}
	// Exibe: [1 2 3 4 5]
	fmt.Println(array3)
}
```

- Usando slice

```go
package main

import "fmt"

func main() {
	// var nome_do_array[tamanho do array] tipo do array
	var array1[5] int
	// Exibe: [0 0 0 0 0]
	fmt.Println(array1)

	// Populando o array
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	// Exibe: [1 2 3 4 5]
	fmt.Println(array1)

	// Podemos declarar arrays usando inferência de tipo
	array2 := [5] string {
		"Posição1", "Posição 2", "Posição 3", "Posição 4", "Posição 5",
	}

	// Exibe: [Posição1 Posição 2 Posição 3 Posição 4 Posição 5]
	fmt.Println(array2)

	// Outra forma de declarar um array é usando o "operador spread" semelhante a do JavaScript que calcula automaticamente o tamanho do array.
	array3 := [...] int {
		1,
		2,
		3,
		4,
		5,
	}
	// Exibe: [1 2 3 4 5]
	fmt.Println(array3)

	// Array slice com 5 valores do tipo string
	slice := [] string {
		"Slice 1",
		"Slice 2",
		"Slice 3",
		"Slice 4",
		"Slice 5",
	}

	// Exibe: [Slice 1 Slice 2 Slice 3 Slice 4 Slice 5]
	fmt.Println(slice)

	// Usando a função append(slice, valor) para adicionar valores no slice. A função retorna um slice novo com o item incluido.
	slice = append(slice, "Slice 6")
	// Exibe: [Slice 1 Slice 2 Slice 3 Slice 4 Slice 5 Slice 6]
	fmt.Println(slice)

	// Usando o slice para "fatiar" o array da posição 1(inclusivo) até a 3(menos a 3 "exclusivo").
	slice2 := array2[1:3]
	// Exibe: [Posição 2 Posição 3]
	fmt.Println(slice2)

	// Alterando o valor do array2
	array2[1] = "Posição alterada"
	// Exibe: [Posição alterada Posição 3]
	fmt.Println(slice2)
}
```

- **Ou seja, podemos criar um slice a partir de um array, e o slice é uma referência do array, ou seja, se alterarmos o valor do array, o slice também será alterado. Ele funciona como um ponteiro. Note que usamos a notação \[1:3\] para indicar que queremos o slice a partir da posição 1 até a posição 3, mas sem incluir a posição 3.**

- **Arrays internos:** Quando criamos um slice, o Go cria um array interno com o tamanho do slice e o slice é uma referência para esse array interno. Quando adicionamos um novo item no slice, o Go cria um novo array interno com o tamanho do slice + 1 e copia os valores do array antigo para o novo array. O slice é atualizado para apontar para o novo array interno.

```go
package main

import "fmt"

func main() {
	// var nome_do_array[tamanho do array] tipo do array
	var array1[5] int
	// Exibe: [0 0 0 0 0]
	fmt.Println(array1)

	// Populando o array
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	// Exibe: [1 2 3 4 5]
	fmt.Println(array1)

	// Podemos declarar arrays usando inferência de tipo
	array2 := [5] string {
		"Posição1", "Posição 2", "Posição 3", "Posição 4", "Posição 5",
	}

	// Exibe: [Posição1 Posição 2 Posição 3 Posição 4 Posição 5]
	fmt.Println(array2)

	// Outra forma de declarar um array é usando o "operador spread" semelhante a do JavaScript que calcula automaticamente o tamanho do array.
	array3 := [...] int {
		1,
		2,
		3,
		4,
		5,
	}
	// Exibe: [1 2 3 4 5]
	fmt.Println(array3)

	// Array slice com 5 valores do tipo string
	slice := [] string {
		"Slice 1",
		"Slice 2",
		"Slice 3",
		"Slice 4",
		"Slice 5",
	}

	// Exibe: [Slice 1 Slice 2 Slice 3 Slice 4 Slice 5]
	fmt.Println(slice)

	// Usando a função append(slice, valor) para adicionar valores no slice. A função retorna um slice novo com o item incluido.
	slice = append(slice, "Slice 6")
	// Exibe: [Slice 1 Slice 2 Slice 3 Slice 4 Slice 5 Slice 6]
	fmt.Println(slice)

	// Usando o slice para "fatiar" o array da posição 1(inclusivo) até a 3(menos a 3 "exclusivo").
	slice2 := array2[1:3]
	// Exibe: [Posição 2 Posição 3]
	fmt.Println(slice2)

	// Alterando o valor do array2
	array2[1] = "Posição alterada"
	// Exibe: [Posição alterada Posição 3]
	fmt.Println(slice2)

	// A função make aloca espaço na memória do computador para uma determinada coisa do nosso programa.
	// A função make possui três parâmetros: []tipo_de_dado, quantidade_de_itens e quantidade_maxima_de_itens
	slice3 := make([]float32, 10, 11)
	// Exibe: [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(slice3)

	// Usando a função len() para saber o total de elementos presentes no array
	fmt.Println(len(slice3)) // Retorna 10

	// Usando a função cap() para saber a capacidade máxima de elementos que o array suporta
	fmt.Println(cap(slice3)) // Retorna 11

	// Estourando a capacidade máxima do slice3
	slice3 = append(slice3, 11)
	// Exibe: [0 0 0 0 0 0 0 0 0 0 11]
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Retorna 11
	fmt.Println(cap(slice3)) // Retorna 11

	// Adicionando mais um elemento no array temos
	slice3 = append(slice3, 12)
	// Exibe: [0 0 0 0 0 0 0 0 0 0 11 12]
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Retorna 12
	// Note que retornou 24. Quando o slice é "estourado", o Go dobra a capacidade máxima do array baseado na quantidade de elementos no array.
	fmt.Println(cap(slice3)) // Retorna 24

	// Caso a gente omita a capacidade máxima, temos que a capacidade máxima será igual ao tamanho do array.
	slice4 := make([]int, 5)
	// Exibe: [0 0 0 0 0]
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Retorna 5
	fmt.Println(cap(slice4)) // Retorna 5

}
```

- Em resumo: A função make() aloca espaço na memória do computador para uma determinada coisa do nosso programa. A função make possui três parâmetros: []tipo_de_dado, quantidade_de_itens e quantidade_maxima_de_itens. A função len() retorna o total de elementos presentes no array. A função cap() retorna a capacidade máxima de elementos que o array suporta. Quando o slice é "estourado", o Go dobra a capacidade máxima do array baseado na quantidade de elementos no array como visto acima.

## Maps

- Maps são estruturas de dados que permitem armazenar dados do tipo chave-valor. A chave pode ser de qualquer tipo, mas o valor deve ser um tipo específico. A sintaxe para declarar um map é:

```go
package main

import "fmt"

func main() {
	// nome_do_map := map[tipo_da_chave]tipo_dos_valores
	usuario := map[string]string {
		// Chave: Valor
		"nome": "Douglas",
		"sobrenome": "Cunha",
	}
	// Exibe: map[nome:Douglas sobrenome:Cunha]
	fmt.Println(usuario)

	// Acessando cada chave de forma individual
	fmt.Println(usuario["nome"]) // Douglas
	fmt.Println(usuario["sobrenome"]) // Cunha


	// Usando maps aninhados
	// Temos a chave do tipo string e o valor é uma estrutura map do tipo string-string
	usuario2 := map[string]map[string]string {
		// Nome é a nosas chave, e o valor é um objeto que é o segundo map.
		"nome": {
			"primeiro-nome": "Douglas",
			"segundo-nome": "João",
			"terceiro-nome": "Matheus",
		},
	}
	// Exibe: map[nome:map[primeiro-nome:Douglas segundo-nome:João terceiro-nome:Matheus]]
	fmt.Println(usuario2)

	// Acessando cada valor de forma individual
	// Usamos a notação de colchetes duplo, onde temos chave-valor.
	fmt.Println(usuario2["nome"]["primeiro-nome"]) // Douglas
	fmt.Println(usuario2["nome"]["segundo-nome"]) // João
	fmt.Println(usuario2["nome"]["terceiro-nome"]) // Matheus

	// Deletando uma chave usando a funão delete(nome-do-map, nome-da-chave)
	delete(usuario2, "nome") // Deleta a chave e os valores.
	fmt.Println(usuario2) // Exibe um map vazio. map[]

	// Adicionando uma nova chave com a notação nome_do_map["nome_da_chave_nova"] = map[tipo_de_dado]tipo_de_dado
	usuario2["aboutme"] = map[string]string {
		"nome": "Douglas",
	}

	fmt.Println(usuario2["aboutme"]["nome"]) // Douglas
}
```

- Em resumo: Maps são estruturas de dados que permitem armazenar dados do tipo chave-valor. A chave pode ser de qualquer tipo, mas o valor deve ser um tipo específico. A sintaxe para declarar um map é: nome_do_map := map[tipo_da_chave]tipo_dos_valores. A função delete(nome-do-map, nome-da-chave) deleta uma chave e seus valores. A notação nome_do_map["nome_da_chave_nova"] = map[tipo_de_dado]tipo_de_dado adiciona uma nova chave com um valor.

## Estrutura de Controle

- Estruturas de controle são utilizadas para controlar o fluxo de execução do programa. No Go, temos as seguintes estruturas de controle: if, else, switch, for, break, continue, goto, defer, panic e recover.

- `Estrutura if/else`

```go
package main

import "fmt"

func main() {
	// Declaramos uma variável idade do tipo int e atribuimos um valor a ela.
	idade := 18

	// Analisa se a condição é verdadeira, caso seja, executa o bloco if, caso contrário executa o bloco else.
	if idade >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade")
	}
}
```

Em Go não é necessário usar parênteses na estrutrua if, é opcional.

- `Declarando uma variável e inicializando usando if init`

```go
package main

import "fmt"

func main() {
	// Declaramos uma variável idade do tipo int e atribuimos um valor a ela.
	idade := 18

	// Analisa se a condição é verdadeira, caso seja, executa o bloco if, caso contrário executa o bloco else.
	if idade >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade")
	}

	// Atribuindo um novo valor a variável idade
	idade = 17

	// Primeiro declaramos uma variável e inicializamos ela, nesse caso com a variável idade.
	// Em seguida usamos ponto e vírgula e fazemos a condição (outraIdade < 18)
	// Caso seja true, executa o bloco abaixo.
	if outraIdade := idade; outraIdade < 18 {
		fmt.Println("Menor de idade!")
	} else {
		fmt.Println("Maior de idade!")
	}

	// Não podemos acessar a variável outraIdade fora da estrutura if.
	// fmt.Println(outraIdade) resulta em undefined, ou seja, variável não definida.
}
```

A estrutura if init serve para declarar e inicializar uma variável dentro da estrutura if. Essa variável só pode ser acessada dentro da estrutura if. Note que a condição é realizada na mesma linha da declaração da variável.

- Usando a estrutura if, else e else if

```go
package main

import "fmt"

func main() {
	// Declaramos uma variável idade do tipo int e atribuimos um valor a ela.
	idade := 18

	// Analisa se a condição é verdadeira, caso seja, executa o bloco if, caso contrário executa o bloco else.
	if idade >= 18 {
		fmt.Println("Maior de idade!")
	} else {
		fmt.Println("Menor de idade")
	}

	// Atribuindo um novo valor a variável idade
	idade = 8

	// Primeiro declaramos uma variável e inicializamos ela, nesse caso com a variável idade.
	// Em seguida usamos ponto e vírgula e fazemos a condição (outraIdade < 18)
	// Caso seja true, executa o bloco abaixo.
	if outraIdade := idade; outraIdade >= 18 {
		fmt.Println("Maior de idade!")
	} else if outraIdade < 10 {
		// Exibe: Você tem menos de 10 anos!
		fmt.Println("Você tem menos de 10 anos!")
	} else {
		fmt.Println("Menor de idade!")
	}
}
```

Ou sejam, a estrutura else if testa para mais de uma condição. Caso a primeira condição seja falsa, ele testa a segunda condição, caso seja falsa, ele testa a terceira condição e assim por diante.

## Estrutura Switch

A estrutura switch testa para uma condição e executa o bloco de código correspondente. A estrutura switch é semelhante a estrutura if, porém, a estrutura switch é mais utilizada quando temos muitas condições a serem testadas.

```go
package main

import "fmt"

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
		// Caso padrão onde o usuário não digita um número entre 1 e 7.
	default:
		return "Número inválido!"
	}
}

// Outra forma de usar a estrutura switch
func diaDaSemana2 (numero int) string {
	// Colocamos a condição no case.
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
		// Caso padrão
	default:
		return "Número inválido!"
	}
}

func main() {

	// Exibe Terça-Feira
	dia3 := diasDaSemana(3)
	fmt.Println(dia3)

	// Exibe Número inválido!
	dia8 := diasDaSemana(8)
	fmt.Println(dia8)

	// Testando para a segunda forma de usar o switch
	// Exibe Quinta-Feira
	dia1 := diaDaSemana2(5)
	fmt.Println(dia1)
	// Exibe Número inválido!
	dia10 := diaDaSemana2(10)
	fmt.Println(dia10)
}
```

Note que em Go não é necessário colocar o break no final de cada case. O Go já entende que quando ele encontrar um case que seja verdadeiro, ele deve executar o bloco de código correspondente e sair da estrutura switch.

## Estruturas de Repetição

As estruturas de repetição são utilizadas para executar um bloco de código várias vezes. Existem três estruturas de repetição em Go: **for e for range**

- Estrutura for

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// Leia: Enquanto i for menor que 10, faça algo.
	for i < 10 {

		// Incrementa a variável i em uma unidade a cada loop.
		i++
		// "Para" o código por um segundo.
		time.Sleep(time.Second)
		fmt.Println("Incrementando i: ", i)
	}

	fmt.Println("Valor final do i: ", i)

	// Usando a notação init temos
	for j := 0; j <= 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
}
```

No exemplo acima usamos a estrutura for, onde passamos uma condição `i < 10` e `j <= 10`. Caso a condição seja verdadeira, o bloco de código é executado, caso contrário, o bloco de código é ignorado. Além disso, usamos a notação init, onde declaramos uma variável e inicializamos ela dentro do for. Essa variável só existe dentro do for, ou seja, não pode ser acessada fora do for.

- Estrutura for range

A estrutura for range é utilizada para percorrer um array, slice, map ou string. Veja o exemplo abaixo:

```go
package main

import (
	"fmt"
)

func main() {
	// A estutura for range é usado quando queremos iterar sobre arrays, strings, maps etc.
	// Declaramos um array do tipo string contendo três strings.
	nomes := [3] string {
		"Douglas",
		"Lucas",
		"Marcel",
	}

	// Iterando sobre o array
	// Declaramos duas variáveis, o índice e o valor do array, em seguida usamos range e o nome do array.
	for index, value := range nomes {
		// Exibe o índice e o valor.
		fmt.Println(index, value)
	}

	// Caso não queiramos o índice, basta usar o underline.
	for _, value := range nomes {
		// Exibe Douglas, Lucas e Marcel.
		fmt.Println(value)
	}

	// Iterando sobre strings
	for index, letra := range "PALAVRA" {
		// Nesse caso exibe o índice e o valor da letra na tabela ASCII.
		fmt.Println(index, letra)
	}

	fmt.Println("---------------------")

	// Para usar a letra no lugar do número da tabela ASCII basta usar a função string().
	// Iterando sobre strings
	for index, letra := range "PALAVRA" {
		// Nesse caso exibe o índice e cada letra durante a execução do loop.
		fmt.Println(index, string(letra))
	}

	// Iterando sobre um map
	usuario := map[string] string {
		"nome": "Douglas",
		"sobrenome": "Cunha",
	}

	// Para iterar sobre um map, primeiro temos uma chave e depois um valor.
	fmt.Println("-------------------------------------")
	for chave, valor := range usuario {
		fmt.Println("Chave:", chave, "\nValor:", valor)
	}
}
```

Em resumo, a estrutura for range é utilizada para percorrer um array, slice, map ou string. Para cada iteração, o for range retorna o índice e o valor do array, slice, map ou string. Caso não queiramos o índice, basta usar o underline. Para usar a letra no lugar do número da tabela ASCII basta usar a função string(). Para iterar sobre um map, primeiro temos uma chave e depois um valor. Por fim, Go não permite iterar sobre uma struct.

## Retorno Nomeado de Funções

Em Go, é possível retornar mais de um valor de uma função. Além disso, podemos dar nomes aos valores de retorno. Veja o exemplo abaixo:

```go
package main

import "fmt"

// Note que no nosso retorno nós passamos dois nomes, soma e subtracao.
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	// Realizamos o cálculo no retorno da função
	soma = n1 + n2
	subtracao = n1 - n2
	// Passamos return no final.
	return
}

func main() {
	// Chamamos a função e armazenamos os seus respectivos valores.
	soma, subtracao := calculosMatematicos(5, 2)
	// Exibe 7 e 3.
	fmt.Println(soma, subtracao)
}
```

Note que no exemplo acima nós passamos dois valores como retorno e dois nomes para os valores de retorno. Em seguida, realizamos o cálculo no retorno da função e passamos return no final. Por fim, chamamos a função e armazenamos os seus respectivos valores. Funções com retorno nomeado deixa nosso código mais legível e organizado.

## Funções Variáticas

Em Go, é possível passar uma quantidade variável de argumentos para uma função. Para isso, basta usar três pontos antes do tipo do argumento. Veja o exemplo abaixo:

```go
package main

import "fmt"

// Passamos um parâmetro seguido de três pontos e o tipo de dado do parâmetro + o tipo de dado de retorno.
func soma(numeros ...int) int {
	// Criamos um contador
	total := 0
	// Usamos o for para percorrer o slice que é gerado pela quantidade de parâmetros.
	for _, numero := range numeros {
		// Somamos o numero no parâmetro com o total.
		total += numero
	}
	// Retornamos o total.
	return total
}

// Função que recebe uma string e uma sequência de números.
func escrever(texto string, numeros ...int) {
	// Percorre o slice numeros e exibe na tela a string e o numero.
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// Chamamos a função passando 11 parâmetros e armazenamos o resultado.
	totalDaSoma := soma(1, 5, 10, 20, 100, 350, 3, 2, 1, 6, 123)

	// Exibe 621.
	fmt.Println(totalDaSoma)

	// Chamamos a função e passamos os dois argumentos.
	escrever("Olá, mundo!", 10, 20, 30, 40, 50)
}
```

Em resumo, funções variáticas são funções que recebem uma quantidade variável de argumentos. Para isso, basta usar três pontos antes do tipo do argumento. Além disso, só podemos ter um parâmetro variático por função.

## Funções Anônimas

Em Go, é possível criar funções sem nome. Essas funções são chamadas de funções anônimas. Veja o exemplo abaixo:

```go
package main

import "fmt"

func main() {
	// Função anônima que exibe Hello, world! na tela.
	func () {
		fmt.Println("Hello, world!")
	}() // É obrigatório o uso do parênteses no final.

	// Usando parâmetros em funções anônimas
	// Declaramos o parâmetro e o tipo e passamos o argumento dentro dos parênteses.
	func (user string) {
		fmt.Println("Olá", user)
	} ("Douglas")

	// Podemos criar uma função anônima que retorna algo e armazenar ela em uma variável.
	retorno := func (texto string) string {
		// A função Sprintf formata a saída, note que passamos %s que significa que queremos exibir o valor do texto que é do tipo string.
		return fmt.Sprintf("Retornando: %s", texto)
	}("Olá, eu sou um texto :)")
	// Exibimos o resultado do retorno.
	fmt.Println(retorno)
}
```

Note que no exemplo acima criamos duas funções anônimas, uma exibe Hello, world! na tela e a outra exibe Olá Douglas. Além disso, criamos uma função anônima que retorna algo e armazenamos ela em uma variável. Em resumo, funções anônimas são funções sem nome. Essas funções são muito úteis quando queremos passar uma função como argumento para outra função.

## Funções Recursivas

Em Go, é possível criar funções recursivas. Funções recursivas são funções que chamam a sí próprio n vezes. Veja o exemplo abaixo:

```go
package main

import "fmt"

// Função recursiva que recebe um inteiro positvo e retorna um inteiro positivo.
func fibonacci (numero uint) uint {
	// Caso base de parada
	if numero <= 1 {
		return numero
	}
	// Caso recursivo
	return fibonacci(numero - 1) + fibonacci(numero - 2)
}

func main() {
	// Armazenamos o valor 5 na variável numero
	var numero uint = 5
	// Chamamos a função fibonacci e passamos numero como argumento
	fmt.Println(fibonacci(numero))

	// Usando a estrutura for para exibir todos os valores da sequência até n valor.
	// passamos uint(0) indicando que queremos que o for comece em 0.
	fmt.Println("---------------------")
	for i := uint(0); i <= numero; i++ {
		fmt.Println(fibonacci(i))
	}
}
}
```

A sequência de Fibonacci é uma sequência de números inteiros, começando normalmente por 0 e 1, na qual, cada termo subsequente corresponde à soma dos dois anteriores. A sequência de Fibonacci é uma sequência infinita. No exemplo acima, criamos uma função recursiva que recebe um inteiro positivo e retorna um inteiro positivo. Em seguida, chamamos a função e passamos o valor 5 como argumento. Por fim, exibimos o resultado na tela.

- Note que em funções recursivas é necessário ter um caso base de parada, caso contrário, a função vai entrar em loop infinito gerando um estouro de memória.

- Funções recursivas geram um grande consumo de memória, por isso, é recomendado que se use funções recursivas apenas quando for necessário.

# Defer

Em Go, é possível usar a palavra reservada defer. A palavra defer faz com que uma **função seja executada no final do escopo da função.** Veja o exemplo abaixo:

```go
package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(notaUm, notaDois int) bool {
	// O defer abaixo é executado antes da estrutura if.
	defer fmt.Println("Média calculada. Resultado será retornado") // Executa em seguida
	fmt.Println("Entrando na função para verificar se o aluno foi aprovado") // Executa primeiro

	media := (notaUm + notaDois) / 2

	// Por fim, retorna true ou false.
	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	// Chamando as duas funções
	funcao1()
	funcao2()


	// DEFER = ADIAR.

	fmt.Println("---------------------")
	// Usando a palavra-reservada defer
	// defer faz com que a execução da função 1 se dê apenas quando a função 2 for executada.
	defer funcao1()
	funcao2()

	fmt.Println("---------------------")

	// Note que primeiro executamos a função alunoEstaReprovado e a funcao1 é executada por último.
	fmt.Println(alunoEstaAprovado(7, 8))
}
```

Note que na função alunoEstaAprovado nós usamos a palavra-reservada defer em um fmt.Println. Isso quer dizer que a função Println será executada antes da estrutura if, isso se deve ao fato da estrutura if possuir retorno. Ou seja, defer "retarda" a execução da função até o final do escopo da função.

## Panic e Recover

Em Go, é possível matar a execução do nosso programa usando a função panic. Para recuperar a execução do programa, usamos a função recover. Veja o exemplo abaixo:

```go
package main

import "fmt"

func recuperarExecucao() {
	// Usando a função recover
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(notaUm, notaDois float64) bool {
	defer recuperarExecucao()
	media := (notaUm + notaDois) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// Faz o sistema entrar em "panico" e mata a execução do programa.
	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
```

Resultado:

```bash
go run panic-recover.go
Execução recuperada com sucesso!
false
Pós execução!
```

A função panic faz com que o programa entre em pânico e seja interrompido. A função recover faz com que o programa seja recuperado e continue a execução. Note que utilizamos um if init para verificar se o valor de r é diferente de nil. Isso se deve ao fato de que a função recover retorna um valor do tipo interface, ou seja, pode retornar qualquer tipo de dado. Por isso, é necessário verificar se o valor retornado é diferente de nil. Para mais informações acesse a [documentação.](https://pkg.go.dev/builtin#recover)

## Função Closure

Em Go, é possível criar funções anônimas. Essas funções anônimas são chamadas de closures. Veja o exemplo abaixo:

```go
package main

import "fmt"

func closure() func() {
	// Variável texto da função closure.
	texto := "Dentro da função closure!"

	// A função closure acessa a variável externa a ela.
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	// Armazenamos o resultado da função e chamamos a funcaoNova.
	funcaoNova := closure()
	funcaoNova()
}
```

Ou seja, closures são funções que podem acessar variáveis externas a elas. Isso é possível porque closures são funções que retornam funções.

## Funções com Ponteiros

Em Go, é possível usar ponteiros em funções. Veja o exemplo abaixo:

```go
package main

import "fmt"

// A função abaixo recebe um ponteiro de int
func inverterSinal(numero *int) {
	// Usamos o ponteiro de int e retornamos o seu valor com o sinal trocado.
	// Note que não usamos o return, pois a alteração é feito diretamente no endereço de memória.
	*numero = *numero * -1
}

func main() {
	numero := 20
	// Chamamos a função numero passando a variável e o local da memoria em que a variável numero está.
	inverterSinal(&numero)
	// Retorna o valor da variável com sinal trocado.
	fmt.Println(numero)
}
```

Basicamente o que fazemos é passar o endereço de memória da variável para a função. Dessa forma, a função pode alterar o valor da variável diretamente no endereço de memória. Dizemos que estamos passando um valor por referência. Note que a função não faz uso da palavra-reservada return, pois ela altera o valor da variável diretamente no endereço de memória.

## Função init

Em Go, é possível criar uma função init. Essa função é executada antes da função main. Veja o exemplo abaixo:

```go
package main

import "fmt"

// A função init é executada antes da função main().
func init() {
	fmt.Println("Executado a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}
```

A função init() é executada antes da função main(). Além disso, só pode existir uma função init() por arquivo e não por pacote.

## Métodos

Em Go, é possível criar métodos. Métodos são funções que pertencem a uma estrutura. Veja o exemplo abaixo:

```go
package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

// Passamos a estrutura usuario que possui o método salvar().
func (u usuario) salvar() {
	// Acessamos a nossa variável nome e idade através do parâmetro u da função.
	fmt.Printf("Nome: %s\nIdade: %d\n", u.nome, u.idade)
}

// Método que retorna se um usuário é maior de idade
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Alterando o valor de uma variável usando ponteiro de struct
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario {
		"Douglas",
		23,
	}

	fmt.Println(usuario1)
	// Chamamos o método salvar().
	usuario1.salvar()
	// Chamamos o método maiorDeIdade que retorna true.
	fmt.Println(usuario1.maiorDeIdade())

	// Chamamos o método fazerAniversario() para alterar o valor da variável.
	usuario1.fazerAniversario()
	// Exibimos a idade do usuário.
	fmt.Println(usuario1.idade) // Exibe 24.
}
}
```

Ou seja, **métodos são funções que pertencem a uma estrutura**. Para criar um método, basta **passar a estrutura como parâmetro da função.** Dessa forma, podemos **acessar os atributos da estrutura.** Além disso, note que podemos **passar a estrutura por referência**. Dessa forma, podemos **alterar o valor da variável diretamente no endereço de memória.** Ou seja, usamos um ponteiro de struct.

## Interfaces

Em Go, é possível criar interfaces. Interfaces são tipos abstratos que definem um conjunto de métodos. Veja o exemplo abaixo:

```go
package main

import (
	"fmt"
	"math"
)

// Uma interface do tipo forma.
type forma interface {
	// Método area() que retorna um valor do tipo float64.
	area() float64
}

// A função escreverArea recebe uma interface e retorna o método area().
// Nota: Para usar o método area() é necessário que o struct retangulo ou circulo tenha um método chamado area() do tipo float64.
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// Criamos uma função area() da struct retangulo.
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// Criamos um método area que recebe uma struct chamada circulo.
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}



func main() {
	r := retangulo {
		10,
		15,
	}

	c := circulo {
		10,
	}

	// Exibimos o resultado.
	escreverArea(r)
	escreverArea(c)
}
```

Ou seja, interfaces armazenam um conjunto de métodos. Dessa forma, podemos criar uma interface e passar uma estrutura que tenha os métodos da interface. Além disso, podemos criar uma função que recebe uma interface como parâmetro. Dessa forma, podemos passar uma estrutura que tenha os métodos da interface.

**NOTA:** `Para usar os métodos da interface é necessário que a estrutura(struct) tenha os métodos da interface, obedecendo o nome e o tipo de retorno do método da interface, como é o caso do método area() acima.`

## Interfaces genéricas

Em Go, é possível criar interfaces genéricas. Interfaces genéricas são interfaces que não possuem métodos. Veja o exemplo abaixo:

```go
package main

import "fmt"

// A função generica recebe um argumento do tipo interface.
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// Note que podemos passar qualquer tipo de dado que será exibido. Por isso o "genérico".
	generica("String") // String
	generica(12) // int
	generica(12.3) // float
	generica(true) // bool
}
```

Ou seja, interfaces genéricas são interfaces que não possuem métodos. Dessa forma, podemos passar qualquer tipo de dado para a função que possui uma interface genérica como parâmetro. Além disso, podemos usar a função `fmt.Println()` que possui uma interface genérica como parâmetro.

[Voltar](../README.md)
