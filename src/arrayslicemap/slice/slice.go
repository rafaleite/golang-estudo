package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [5]int{1, 2, 3, 4, 5} // array
	s1 := []int{1, 2, 3, 4, 5}  // slice

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))
	fmt.Println(cap(s1))

	// pode-se imaginar slice como: tamanho e ponteiro para um elemento de um array
	s2 := a1[:5]
	a1[0] = 10
	s3 := append(s2, 7, 8, 9, 10, 11, 12)
	a1[1] = 11

	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(a1)
	fmt.Println(s3)

	//=======================================
	s5 := make([]int, 10, 20)
	fmt.Println(s5, len(s5), cap(s5))
}
