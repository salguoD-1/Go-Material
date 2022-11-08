package main

import (
  "fmt"
)

// Estrutura do tipo usuario
type usuario struct {
  nome string
  idade uint8
  email string
  endereco endereco
}

// Estrutura do tipo endereco
type endereco struct {
  logradouro string
  numero uint8
}

// Função principal
func main() {
  // Declaramos uma variável chamada user do tipo usuario e atribuímos valores a cada uma das variáveis da estrutura.
  user := usuario{
    nome: "Douglas",
    idade: 23,
    email: "douglas@email.com",
    // Atribuímos valores a cada uma das variáveis da estrutura endereco.
    endereco: endereco{
      logradouro: "Rua aprendendo Go",
      numero: 13,
    },
  }

  // Exibe o resultado de cada variável
  fmt.Println(user.nome) // Douglas
  fmt.Println(user.idade) // 23
  fmt.Println(user.email) // douglas@email.com
  fmt.Println(user.endereco.logradouro) // Rua aprendendo Go
  fmt.Println(user.endereco.numero) // 13
}