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