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