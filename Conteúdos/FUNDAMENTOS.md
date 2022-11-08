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

- **Arrays internos**

[Voltar](../README.md)
