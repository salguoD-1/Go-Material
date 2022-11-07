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