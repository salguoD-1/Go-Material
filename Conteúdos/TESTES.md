# Testes Automatizados

## O que são testes automatizados?

Testes automatizados são testes de software que são executados automaticamente. Os testes automatizados são executados por um software de teste, que é um programa que controla a execução dos testes. Os testes automatizados são executados repetidamente, em intervalos regulares, para verificar se o software está funcionando corretamente.

## Por que testar?

Testes são importantes para garantir a segurança que as aplicações precisam. Os testes automatizados são uma forma de garantir que o software está funcionando corretamente, sem a necessidade de intervenção humana.

## Como testar?

Existem várias formas de testar um software. Os testes automatizados são uma delas. Os testes automatizados são feitos por um software de teste, que é um programa que controla a execução dos testes. Os testes automatizados são executados repetidamente, em intervalos regulares, para verificar se o software está funcionando corretamente.

A linguagem Go provem uma biblioteca padrão para testes, que é a `testing`. Essa biblioteca fornece uma estrutura para escrever testes e executá-los. Os testes são escritos em arquivos com o nome `*_test.go`, que são arquivos Go que contém testes. Os testes são escritos em funções que começam com `Test`, seguido do nome da função. Por exemplo, a função `TestHelloWorld` é um teste que testa a função `HelloWorld`.

### Testes Unitários

Testes Unitários são testes que testam uma única função.

Para usar testes unitários, é necessário importar a biblioteca `testing` e criar um arquivo que possua no nome \_test.go. Além disso, a função de teste deve conter o nome Test antes do nome da função, como visto abaixo:

```go
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
```

Note que a função possui Test no nome, em seguida passamos um ponteiro para o tipo `testing.T` como parâmetro. Esse ponteiro é usado para exibir erros no terminal. A função `t.Error` exibe um erro no terminal. Por fim, compara o valor recebido com o valor esperado.

Para executar o teste, basta entrar na pasta que se encontra o seu arquivo e o teste e digitar:

```go
go test
PASS
ok      testes/enderecos        0.002s
```

Note que retornou PASS, isso quer dizer que a função está funcionando perfeitamente.

Caso teste falhasse, teriamos:

```go
go test
--- FAIL: TestTipoDeEndereco (0.00s)
    enderecos_test.go:15: O tipo recebido é diferente do esperado!
FAIL
exit status 1
FAIL    testes/enderecos        0.001s
```

Note que exibe o erro no terminal.

[Voltar](../README.md)
