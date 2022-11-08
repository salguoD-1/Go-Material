package main

import "fmt"

// Struct do tipo pessoa
type pessoa struct {
	nome string
	sobrenome string
	idade int8
	altura uint8
	peso float32
}

// Struct do tipo estudante
type estudante struct {
	// Passamos o nome da nossa struct que queremos "herdar". Não é necessário passar o tipo.
	pessoa
	curso string
	faculdade string
}

func main() {
	// Declaramos uma variável chamada pessoaUm do tipo pessoa e setamos alguns valores.
	pessoaUm := pessoa {
		"Douglas",
		"Cunha de Jesus",
		23,
		172,
		68.5,
	}

	// Exibe: {Douglas Cunha de Jesus 23 172 68.5}
	fmt.Println(pessoaUm)

	// Declaramos uma variável chamada estudanteUm do tipo estudante e passamos uma struct(pessoaUm) e setamos dois valores.
	estudanteUm := estudante {
		pessoaUm,
		"Sistemas de Informação",
		"Universidade Federal de Sergipe",
	}

	// Exibe: {{Douglas Cunha de Jesus 23 172 68.5} Sistemas de Informação Universidade Federal de Sergipe}
	fmt.Println(estudanteUm)

	// Para acessar as variáveis da struct pessoa basta usar a notação dot(.) seguido do nome da variável.

	fmt.Println(estudanteUm.nome)
	fmt.Println(estudanteUm.sobrenome)
	fmt.Println(estudanteUm.idade)
	fmt.Println(estudanteUm.altura)
	fmt.Println(estudanteUm.peso)
	fmt.Println(estudanteUm.curso)
	fmt.Println(estudanteUm.faculdade)
}