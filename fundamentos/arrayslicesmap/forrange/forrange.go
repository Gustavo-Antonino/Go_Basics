package main

import "fmt"

func main() {
	// Declaração de um array de inteiros chamado "numeros", sem esses 3 pontons dentro das colchetes "[...]" ele vira um um slice, para falar que é um array, sempre coloque 3 pontos.
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	// Loop para iterar sobre o array "numeros"
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero) // Imprime o índice (incrementado em 1) e o valor do elemento
	}

	// Loop para imprimir apenas os valores do array
	for _, num := range numeros {
		fmt.Println(num)
	}
}
