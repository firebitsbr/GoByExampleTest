// Go tem vários tipos de valor,
// incluindo strings, inteiros, flutuantes, booleanos, etc.
// Aqui estão alguns exemplos básicos.

package main

import "fmt"

func main() {

	//Strings, que podem ser adicionados junto com +
	fmt.Println("go" + "lang")

	//Inteiros e flutuantes.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	//Booleanos, com operadores booleanos como seria de esperar.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
