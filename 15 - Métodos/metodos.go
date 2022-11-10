package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

// Passamos a estrutura usuario que possui o método salvar().
func (u usuario) salvar() {
	// Acessamos a nossa variável nome e idade através do parâmetro u da função.
	fmt.Printf("Nome: %s\nIdade: %d\n", u.nome, u.idade)
}

// Método que retorna se um usuário é maior de idade
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Alterando o valor de uma variável usando ponteiro de struct
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario {
		"Douglas",
		23,
	}

	fmt.Println(usuario1)
	// Chamamos o método salvar().
	usuario1.salvar()
	// Chamamos o método maiorDeIdade que retorna true.
	fmt.Println(usuario1.maiorDeIdade())

	// Chamamos o método fazerAniversario() para alterar o valor da variável.
	usuario1.fazerAniversario()
	// Exibimos a idade do usuário.
	fmt.Println(usuario1.idade) // Exibe 24.
}