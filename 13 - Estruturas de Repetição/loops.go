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