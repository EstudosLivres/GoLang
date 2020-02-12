package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// inteiros
	fmt.Println(1,2,100)
	fmt.Println("Inteiro literal:", reflect.TypeOf(32000))

	// ints sem sinal: uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	// int com sinal
	i1 := math.MaxInt64
	fmt.Println("Valor max int:",i1)
	fmt.Println("o tipo int:",reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("rune é",reflect.TypeOf(i2))
	fmt.Println("rune é",i2)

	var x float32 = 49.99
	fmt.Println("x é do tipo:",reflect.TypeOf(x))
	fmt.Println("49.99 é tipo:",reflect.TypeOf(49.99))

	bo := true
	fmt.Println("bo é do tipo:",reflect.TypeOf(bo))
	fmt.Println("negacao bo:",!bo)

	s1 := "Olá meu nome"
	fmt.Println(s1 + " 1")
	fmt.Println("o tamanho da string é:",len(s1))

	s2 := `Olá
	meu
	nome
	`
	fmt.Println("tamanho s2:",len(s2))

	//var x char = 'a' o tipo char não existe no Go
	char := 'a'
	fmt.Println("tipo char:",reflect.TypeOf(char))
	fmt.Println("valor de char (é o valor unicode):",char)
}
