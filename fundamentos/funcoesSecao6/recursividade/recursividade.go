package main

import "fmt"

func fatorail(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorailAnterior, _ := fatorail(n - 1)
		return n * fatorailAnterior, nil
	}
}

func main() {
	resultado, _ := fatorail(5)
	fmt.Println(resultado)

	_, err := fatorail(-4)
	if err != nil {
		fmt.Println(err)
	}

	// Uma solução melhor seria... uint!

}
