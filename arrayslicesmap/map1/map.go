package main

import "fmt"

func main() {
	// Criando um mapa chamado "aprovados" para associar CPFs a nomes.
	// Maps devem ser inicializados antes de serem usados.
	aprovados := make(map[int]string)

	// Adicionando alguns registros ao mapa.
	aprovados[12345678978] = "Maria"
	aprovados[98765432148] = "Pedro"
	aprovados[90123467890] = "Ana"

	// Imprimindo o mapa completo.
	fmt.Println(aprovados)

	// Iterando sobre o mapa e imprimindo os nomes e CPFs.
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Imprimindo o nome associado ao CPF 90123467890.
	fmt.Println(aprovados[90123467890])

	// Removendo o registro associado ao CPF 90123467890.
	delete(aprovados, 90123467890)

	// Tentando imprimir o nome associado ao CPF 90123467890 novamente.
	// Deve resultar em uma string vazia, pois o registro foi excluído.
	fmt.Println(aprovados[90123467890])
}
