package main

import "fmt"

func fatorail(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorail(n-1)
	}
}

func main() {
	resultado := fatorail(5)
	fmt.Println(resultado)

	// Uma solução melhor seria... uint!

}
