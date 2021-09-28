package main

import "fmt"

func main() {
	fmt.Printf("Média: %.1f\n", media(9.0, 8.4, 6.4, 3.2))
}

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}
