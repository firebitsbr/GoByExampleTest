//for é a única construção de loop de Go. Aqui estão alguns tipos básicos de loops for.
package main

import "fmt"

func main() {

	//O tipo mais básico, com uma única condição.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Um loop clássico inicial/condição/posterior.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//pois sem uma condição fará um loop repetidamente até que você saia do loop ou volte da função envolvente.//Por isso a necessidade do comando break para interruper.
	for {
		fmt.Println("Loop")
		break
	}

	//Você também pode continuar para a próxima iteração do loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
