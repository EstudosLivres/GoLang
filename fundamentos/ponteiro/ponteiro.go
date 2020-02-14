package main

import "fmt"

func main() {
	i := 1

	var p *int = nil //ponteiro apontando para nada
	p = &i // pegando o endereço da variável e alocando no ponteiro p
	*p++   // desreferenciando (pegando o valor) usando o * e incrementando o valor dela, assim sendo esta começou em 1, aqui foi para 2 e na linha de baixo para 3
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
