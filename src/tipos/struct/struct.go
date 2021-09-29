package main

import "fmt"

type Product struct {
	id    uint
	name  string
	price float64
}

type Item struct {
	product  Product
	discount float64
	qty      int
}

type Cart struct {
	userID int
	itens  []Item
}

func initCart() Cart {
	cart := Cart{
		userID: 1,
		itens:  []Item{},
	}

	return cart
}

func (it *Item) getPriceWithDiscount() float64 {
	return it.product.price * (1 - it.discount)
}

func (it *Item) getTotalItemAmount() float64 {
	return it.getPriceWithDiscount() * float64(it.qty)
}

func (c *Cart) addItem(item Item) {
	c.itens = append(c.itens, item)
}

func (c *Cart) printCheckOut() {
	total := 0.0
	qty := 0
	fmt.Println("=============================")
	for i, item := range c.itens {
		total += item.getTotalItemAmount()
		qty += item.qty
		fmt.Printf("%d) %s | Qtde. %d | R$%.2f | total: R$%.2f (%.1f%% de desconto)\n", i+1, item.product.name, item.qty,
			item.product.price, item.getTotalItemAmount(), item.discount*100)
	}
	fmt.Printf("Total: %.2f [%d] itens\n", total, qty)
	fmt.Println("=============================")
}

func main() {
	cart := initCart()
	cart.addItem(Item{
		product:  Product{1, "Cerveja Devassa", 2.25},
		discount: 0.0,
		qty:      12})

	cart.addItem(Item{
		product:  Product{21, "Heineken", 4.00},
		discount: 0.5,
		qty:      24})

	cart.printCheckOut()
}
