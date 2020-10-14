//Em Go, as variáveis são declaradas explicitamente e usadas pelo compilador para, por exemplo,
//verifique a exatidão do tipo das chamadas de função.

package main

import "fmt"

func main() {

	//var declara 1 ou mais vari]áveis.
	var a = "initial"
	fmt.Println(a)

	//Você pode declarar várias variáveis de uma vez.
	var b, c int = 1, 2
	fmt.Println(b, c)

	//Go irá inferir o tipo de variáveis inicializadas.
	var d = true
	fmt.Println(d)

	//Variáveis declaradas sem uma inicialização correspondente têm valor zero.
	//Por exemplo, o valor zero para um int é 0.
	var e int
	fmt.Println(e)

	//A sintaxe: = é uma abreviação para declarar e inicializar uma variável,
	//por exemplo, para var f string = "apple" neste caso.
	f := "apple"
	fmt.Println(f)

}
