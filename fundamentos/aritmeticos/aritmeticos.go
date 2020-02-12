package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Módulo/Resto => ", a%b)

	//bitwise
	fmt.Println("E => ", a&b) // 11 & 10 = 10
	fmt.Println("OU => ", a|b) // 11 | 10 = 11
	fmt.Println("XOR => ", a^b) // 11 ^ 10 = 01

	c := 5.0
	d := 7.0

	//outras operações usando pacote math
	fmt.Println("Maior(a,b) => ", math.Max(float64(a),float64(b)))
	fmt.Println("Menor(c,d) => ", math.Min(c,d))
	fmt.Println("Exponenciacao(c,d) => ", math.Pow(c,d))
}
