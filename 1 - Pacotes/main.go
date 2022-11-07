package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()

	// Retorna nil que significa null. Se não for null, significa que o email é inválido.
	err := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(err)
}
