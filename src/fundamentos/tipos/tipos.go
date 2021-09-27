package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//tipos numericos inteiros
	fmt.Println(reflect.TypeOf(23))

	var b byte = 255
	fmt.Println("O tipo byte é ", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O maximo valor int é", i1)
	fmt.Println("O tipo i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("O valor de i2 é", i2)
	fmt.Println("O tipo de i2 é ", reflect.TypeOf(i2))

	var i3 float32 = 49.99
	fmt.Println("O tipo de i3 é ", reflect.TypeOf(i3))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	var z1 int
	var z2 float64
	var z3 string
	var z4 *int
	var z5 bool

	fmt.Printf("\n%v %v %q %v %v", z1, z2, z3, z4, z5)
}
