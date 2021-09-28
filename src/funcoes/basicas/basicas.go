package main

import "fmt"

func main() {
	nome, idade := multReturn()
	texto := apresentacao(nome, idade)

	fmt.Println(texto)

}

func multReturn() (string, int) {
	return "Rafael", 34
}

func apresentacao(nome string, idade int) string {
	return fmt.Sprintf("Meu nome é %s e tenho %d\n", nome, idade)
}
