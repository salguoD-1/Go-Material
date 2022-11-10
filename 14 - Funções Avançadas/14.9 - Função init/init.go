package main

import "fmt"

// A função init é executada antes da função main().
func init() {
	fmt.Println("Executado a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}