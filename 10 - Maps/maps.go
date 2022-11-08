package main

import "fmt"

func main() {
	// nome_do_map := map[tipo_da_chave]tipo_dos_valores
	usuario := map[string]string {
		// Chave: Valor
		"nome": "Douglas",
		"sobrenome": "Cunha",
	}
	// Exibe: map[nome:Douglas sobrenome:Cunha]
	fmt.Println(usuario)

	// Acessando cada chave de forma individual
	fmt.Println(usuario["nome"]) // Douglas
	fmt.Println(usuario["sobrenome"]) // Cunha


	// Usando maps aninhados
	// Temos a chave do tipo string e o valor é uma estrutura map do tipo string-string
	usuario2 := map[string]map[string]string {
		// Nome é a nosas chave, e o valor é um objeto que é o segundo map.
		"nome": {
			"primeiro-nome": "Douglas",
			"segundo-nome": "João",
			"terceiro-nome": "Matheus",
		},
	}
	// Exibe: map[nome:map[primeiro-nome:Douglas segundo-nome:João terceiro-nome:Matheus]]
	fmt.Println(usuario2)

	// Acessando cada valor de forma individual
	// Usamos a notação de colchetes duplo, onde temos chave-valor.
	fmt.Println(usuario2["nome"]["primeiro-nome"]) // Douglas
	fmt.Println(usuario2["nome"]["segundo-nome"]) // João
	fmt.Println(usuario2["nome"]["terceiro-nome"]) // Matheus

	// Deletando uma chave usando a funão delete(nome-do-map, nome-da-chave)
	delete(usuario2, "nome") // Deleta a chave e os valores.
	fmt.Println(usuario2) // Exibe um map vazio. map[]

	// Adicionando uma nova chave com a notação nome_do_map["nome_da_chave_nova"] = map[tipo_de_dado]tipo_de_dado
	usuario2["aboutme"] = map[string]string {
		"nome": "Douglas",
	}

	fmt.Println(usuario2["aboutme"]["nome"]) // Douglas
}