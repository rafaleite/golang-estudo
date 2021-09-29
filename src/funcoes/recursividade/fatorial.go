package main

import "fmt"

func fatorialSimples(n uint) uint {
	if n == 0 {
		return 1
	}
	anterior := fatorialSimples(n - 1)
	return n * anterior
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		anterior, _ := fatorial(n - 1)
		return n * anterior, nil
	}
}

func main() {
	resultado, _ := fatorial(10)
	fmt.Println(resultado)

	fmt.Println(fatorialSimples(10))
}
