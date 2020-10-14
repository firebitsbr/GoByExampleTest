//Go oferece suporte a constantes de caracteres, string,
//booleanos e valores numéricos.

package main

import (
	"fmt"
)

//const declara um valor constante.
const s string = "constant"

func main() {

	//exibe conteudo em s
	fmt.Println(s)

	//Uma instrução const pode aparecer em qualquer
	//lugar que uma instrução var possa.
	const n = 500000000

	//Expressões constantes realizam
	//operações aritméticas com precisão arbitrária.
	const d = 3e20 / n

	//exibe conteudo em d
	fmt.Println(d)

	//Uma constante numérica não tem tipo
	//até que seja fornecida, como por meio de uma
	//conversão explícita.
	fmt.Println(int64(d))

	//Um número pode receber um tipo usando-o em um
	//contexto que exige um, como uma atribuição de
	//variável ou chamada de função.
	//Por exemplo, aqui math.Sin espera um float64.

}
