package main

import "fmt"

func main() {
	// Cria um mapa chamado funcsPorLetra com letras como chaves e mapas internos como valores
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	// Remove o mapa interno associado à chave "P"
	delete(funcsPorLetra, "P")

	// Itera sobre o mapa funcsPorLetra e imprime as letras e os mapas internos
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
