package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float inferido, pois já recebeu valor

	//forma reduzida de criar var (com o : significa que a var ainda n existe, sem o : significa que ela ja existe)
	area := PI * m.Pow(raio, 2) //se definir a var e n usar o Go dá erro de compilação
	fmt.Println("A área da circunferencia é", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a,b,c,d)

	var e, f bool = true, false
	fmt.Println(e,f)

	g,h,i := 2,false,"pa!"
	fmt.Println(g,h,i)
}
