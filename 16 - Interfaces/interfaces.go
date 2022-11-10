package main

import (
	"fmt"
	"math"
)

// Uma interface do tipo forma.
type forma interface {
	// Método area() que retorna um valor do tipo float64.
	area() float64
}

// A função escreverArea recebe uma interface e retorna o método area().
// Nota: Para usar o método area() é necessário que o struct retangulo ou circulo tenha um método chamado area() do tipo float64.
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// Criamos uma função area() da struct retangulo.
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// Criamos um método area que recebe uma struct chamada circulo.
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}



func main() {
	r := retangulo {
		10,
		15,
	}

	c := circulo {
		10,
	}

	// Exibimos o resultado.
	escreverArea(r)
	escreverArea(c)
}