package main

import "fmt"

type Carro struct {
	modelo string
	ano int
}

type Ferrari struct {
	Carro
	possuiTurbo bool
}

func main() {
	car := Ferrari{}
	car.modelo = "F40"
	car.ano = 2000
	car.possuiTurbo = false

	fmt.Println(car.modelo)
}