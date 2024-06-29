package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10}, // Dessa forma é melhor.
			item{2, 1, 23.49},                         // dessa forma fica meio perdido o que é cada valor.
			item{11, 1000, 3.13},
		},
	}

	fmt.Println("Valor total do pedido é %.2f", pedido.valorTotal())
}
