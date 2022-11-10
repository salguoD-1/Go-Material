package main

import "fmt"

// A função generica recebe um argumento do tipo interface.
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// Note que podemos passar qualquer tipo de dado que será exibido. Por isso o "genérico".
	generica("String") // String
	generica(12) // int
	generica(12.3) // float
	generica(true) // bool
}