package main

import "fmt"

type Cart struct {
	products []Product
}

func (i *Cart) AddProduct(product Product) {
	f := false
	for _, p := range i.products {
		if p.id == product.id {
			p.quantity += product.quantity
			f = true
		}
	}
	if !f {
		i.products = append(i.products, product)
	}
}

func (i *Cart) removeProduct(product Product) {
	for ind, p := range i.products {
		if p.id == product.id {
			p.quantity -= product.quantity
			if p.quantity < 0 {
				i.products = append(i.products[:ind], i.products[ind+1:]...)
			}
		}
	}
}

func (i *Cart) showCart() {
	fmt.Println("******** Added Products in Cart *************** \n")
	for _, p := range i.products {
		fmt.Println(p.name, p.quantity)
	}
}

func (i *Cart) getCartTotal() int {
	cost := 0
	for _, item := range i.products {
		cost += item.quantity * item.price
	}
	return cost
}
