package main

import "fmt"

func main() {

	i := 1

	var p *int = nil

	p = &i // pegando o enderço da variável
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não te aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
