package main

import (
	"fmt"
	"time"
)

func notaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	default:
		return "D"
	}
}

func notaConceitoRefactor(n float64) string {
	var nota = int(n)
	switch {
	case nota >= 9:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 7 && nota < 8:
		return "C"
	default:
		return "D"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Funcao"
	default:
		return "Nao sei"
	}
}

func main() {
	fmt.Println(notaConceitoRefactor(2))
	t := time.Now()
	fmt.Println(t.Format("02/01/2006"))
	fmt.Println(tipo(2.1))
	fmt.Println(tipo(func(){}))
}
