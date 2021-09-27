package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha")

	fmt.Println("Nova linha")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.99
	c := false
	d := "ola"

	fmt.Printf("\n%d %f %.2f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
