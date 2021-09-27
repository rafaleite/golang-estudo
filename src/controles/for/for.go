package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("============== FOR =================")

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}

	for {
		fmt.Println("Para Sempre")
		time.Sleep(time.Second * 2)
	}

}
