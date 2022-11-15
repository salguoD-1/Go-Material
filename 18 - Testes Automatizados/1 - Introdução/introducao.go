package main

import (
	"fmt"
	"testes/enderecos"
)


func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	// Retorna avenida
	fmt.Println(tipoEndereco)
}