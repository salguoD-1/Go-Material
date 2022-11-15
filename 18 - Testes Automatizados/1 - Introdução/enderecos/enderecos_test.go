package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// String para teste
	enderecoParaTeste := "Avenida Paulista"
	// Valor Esperado
	tipoDeEnderecoEsperado := "Rua"
	// Armazena o resultado da função
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		// Exibe um error no terminal
		t.Errorf("O tipo recebido é diferente do esperado!\nEra esperado %s e foi recebido %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}