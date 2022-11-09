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