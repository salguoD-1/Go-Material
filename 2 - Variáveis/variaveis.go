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

	// Constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)
}