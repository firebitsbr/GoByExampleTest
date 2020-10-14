//Ramificar com if e else no Go é simples.

package main

import "fmt"

func main() {

	//Aqui está um exemplo básico de if/else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//Você pode ter uma instrução if sem outro if.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//Uma instrução pode preceder as condicionais; quaisquer variáveis
	//declaradas nesta instrução estão disponíveis em todos os ramos.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

//Observe que você não precisa de parênteses em torno das condições em Go,
//mas as chaves são obrigatórias.

//Não há if ternário (https://en.wikipedia.org/wiki/%3F:) em Go, então você precisará
//usar uma declaração if completa, mesmo para condições básicas.
