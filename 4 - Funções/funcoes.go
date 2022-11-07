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
	// Exibe: 25 -5
	fmt.Println(resultadoSoma)
}