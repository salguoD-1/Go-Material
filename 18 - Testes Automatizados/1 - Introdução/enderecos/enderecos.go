package enderecos

import (
	"strings"
)

// TipoDeEndereco analisa se o tipo de endereço é válido e retorna true ou "Tipo Inválido"
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string {
		"rua",
		"avenida",
		"rodovia",
		"estrada",
	}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	// A função split vai separar a string em um slice contendo fatias da string.
	// Passamos o índice 0 que é o que queremos.
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	// Setamos como false por padrão
	enderecoTemUmTipoValido := false

	// Percorremos o slice e conferimos se o tipo da primeira palavra bate com os tipos do nosso slice.
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	// Retorna a primeira palavra do endereço
	if enderecoTemUmTipoValido {
		// Retorna a primeira palavra em maiúsculo
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo inválido"
}