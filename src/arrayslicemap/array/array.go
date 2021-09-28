package main

import "fmt"

func main() {
	//var notas [3]float64
	//fmt.Println(notas)
	//notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	notas := [...]float64{7.8, 4.3, 9.1}
	fmt.Println(notas)

	total := 0.0

	for _, s := range notas {
		total += s
	}
	fmt.Println(total)

	media := total / float64(len(notas))
	fmt.Printf("Média é %.2f\n", media)
}
