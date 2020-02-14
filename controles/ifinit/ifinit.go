package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// esta variável só será visível e válida dentro da estrutura do if e do else
	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
