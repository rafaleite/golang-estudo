package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 20.4
	b := 2

	fmt.Println(a / float64(b))

	nota := 7.2
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//int para string
	fmt.Println("Teste " + strconv.Itoa(321))

	//string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	c, _ := strconv.ParseBool("true")

	if c {
		fmt.Println("verdadeiro")
	}
}
