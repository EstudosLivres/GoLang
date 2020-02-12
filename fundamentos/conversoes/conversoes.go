package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//CUIDADO, 123 em asc significa '{', 97 significa a
	fmt.Println("Teste " + string(97))

	//int p/ string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string p/ int
	num, _ := strconv.Atoi("123") // o _ significa ignorar caso n usado
	//num, err := strconv.Atoi("123") // se usar assim tem que tratar a var err se n o go n vai compilar
	fmt.Println(num - 122)

	// só vai dar true se passar 1 ou true, qualquer outra coisa é false
	b,_ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}