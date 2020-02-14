package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 5
	i += 4
	i -= 1
	i /= 2
	i *= 2
	i %= 2

	fmt.Println(i)

	x,y := 1,2
	fmt.Println(x,y)

	x,y = y,x
	fmt.Println(x,y)
}
