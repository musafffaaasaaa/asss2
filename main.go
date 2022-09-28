package main

import "fmt"

type Icecream interface {
	getPrice() int
}
type iseCream struct {
}

func (i *iseCream) getPrice() int {
	return 200
}

type chocolateTopping struct {
	cream Icecream
}

func (c *chocolateTopping) getPrice() int {
	creamPrice := c.cream.getPrice()
	return creamPrice + 100
}

type walnutTopping struct {
	cream Icecream
}

func (w *walnutTopping) getPrice() int {
	creamPrice := w.cream.getPrice()
	return creamPrice + 125
}

func main() {
	cream := &iseCream{}

	isecreamWithChocolate := &chocolateTopping{
		cream: cream,
	}

	isecreamWithChocolateAndWalnut := &walnutTopping{
		cream: isecreamWithChocolate,
	}
	fmt.Println("Price of ice cream with chocolate and walnut topping is: ", isecreamWithChocolateAndWalnut.getPrice(), "Tenge")
}
