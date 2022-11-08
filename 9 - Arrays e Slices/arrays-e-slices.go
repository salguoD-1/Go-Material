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