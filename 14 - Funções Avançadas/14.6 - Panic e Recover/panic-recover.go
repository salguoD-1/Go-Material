package main

import "fmt"

func recuperarExecucao() {
	// Usando a função recover
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(notaUm, notaDois float64) bool {
	defer recuperarExecucao()
	media := (notaUm + notaDois) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// Faz o sistema entrar em "panico" e mata a execução do programa.
	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}