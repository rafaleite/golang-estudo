package main

import (
	"fmt"
	"math"
)

func main() {
	PI := 3.1415
	raio := 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println("A área é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 1, "Rafael", false

	fmt.Println(g, h, i)

}
