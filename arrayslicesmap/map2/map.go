package main

import "fmt"

func main() {
	// Cria um mapa chamado funcESalarios com nomes de funcionários e seus salários
	funcESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	// Adiciona um novo funcionário e seu salário ao mapa
	funcESalarios["Rafael Filho"] = 1350.0

	// Remove um funcionário inexistente do mapa
	delete(funcESalarios, "Inexistente")

	// Itera sobre o mapa e imprime os nomes e salários dos funcionários
	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}

	// Itera sobre o mapa e imprime apenas os salários dos funcionários
	for _, salario := range funcESalarios {
		fmt.Println(salario)
	}

	// Itera sobre o mapa e imprime apenas os nomes dos funcionários
	for nome, _ := range funcESalarios {
		fmt.Println(nome)
	}
}
