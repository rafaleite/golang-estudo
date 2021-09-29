package main

import "fmt"

func teste(i *int) {
	*i++
}


func main() {
	n := 2;
	teste(&n)
	fmt.Println(n)
}
