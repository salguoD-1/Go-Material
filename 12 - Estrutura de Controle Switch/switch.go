package main

import "fmt"

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
		// Caso padrão onde o usuário não digita um número entre 1 e 7.
	default:
		return "Número inválido!"
	}
}

// Outra forma de usar a estrutura switch
func diaDaSemana2 (numero int) string {
	// Colocamos a condição no case.
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
		// Caso padrão
	default:
		return "Número inválido!"
	}
}

func main() {

	// Exibe Terça-Feira
	dia3 := diasDaSemana(3)
	fmt.Println(dia3)

	// Exibe Número inválido!
	dia8 := diasDaSemana(8)
	fmt.Println(dia8)

	// Testando para a segunda forma de usar o switch
	// Exibe Quinta-Feira
	dia1 := diaDaSemana2(5)
	fmt.Println(dia1)
	// Exibe Número inválido!
	dia10 := diaDaSemana2(10)
	fmt.Println(dia10)
}