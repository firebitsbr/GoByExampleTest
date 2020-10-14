//As instruções switch expressam condicionais em muitos ramos.

package main

import (
	"fmt"
	"time"
)

func main() {

	//Aqui está uma opção básica.
	i := 2
	fmt.Println("Write ", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//Você pode usar vírgulas para separar várias expressões na
	//mesma instrução case. Usamos o caso padrão opcional neste exemplo também.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch sem uma expressão é uma maneira alternativa de expressar a
	//lógica if/else. Aqui também mostramos como as expressões case
	//podem ser não constantes.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//Um switch de tipo compara tipos em vez de valores.
	//Você pode usar isso para descobrir o tipo de um valor de interface.
	//Neste exemplo, a variável t terá o tipo correspondente a sua cláusula.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("Don't know type \n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
