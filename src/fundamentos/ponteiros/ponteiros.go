package main

func main() {
	a := 2
	b := &a
	var c *int = nil
	c = b
	a++
	print(a, *b, *c)
}
