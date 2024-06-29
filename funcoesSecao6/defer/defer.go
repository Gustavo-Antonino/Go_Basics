package main

import "fmt"

func obterValorAporvado(numero int) int {
	defer fmt.Println("Fim!")
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAporvado(6000))
	fmt.Println(obterValorAporvado(3000))
}
