package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix, em go só existe a forma pós fixada, exemplo: x++, não existe forma pré fixada, exemplo; x--
	x++ // x += 1 ou x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y -1
	fmt.Println(y)

	// EXEMPLO DO QUE NÃO PODE: fmt.Println(x == y--)
}
