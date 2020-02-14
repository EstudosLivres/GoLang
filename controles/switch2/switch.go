package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//primeiro case que retornar true ele vai executar
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite.")
	}
}
