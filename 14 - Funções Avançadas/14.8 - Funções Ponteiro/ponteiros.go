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