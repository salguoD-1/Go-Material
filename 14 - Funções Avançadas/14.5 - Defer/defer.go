package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(notaUm, notaDois int) bool {
	// O defer abaixo é executado antes da estrutura if.
	defer fmt.Println("Média calculada. Resultado será retornado") // Executa em seguida
	fmt.Println("Entrando na função para verificar se o aluno foi aprovado") // Executa primeiro

	media := (notaUm + notaDois) / 2

	// Por fim, retorna true ou false.
	if media >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	// Chamando as duas funções
	funcao1()
	funcao2()


	// DEFER = ADIAR.

	fmt.Println("---------------------")
	// Usando a palavra-reservada defer
	// defer faz com que a execução da função 1 se dê apenas quando a função 2 for executada.
	defer funcao1()
	funcao2()

	fmt.Println("---------------------")

	// Note que primeiro executamos a função alunoEstaReprovado e a funcao1 é executada por último.
	fmt.Println(alunoEstaAprovado(7, 8))
}