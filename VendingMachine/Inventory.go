package main

import "fmt"

type Inventory struct {
	products []Product
}

func (i *Inventory) addProduct(product Product) {
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

func (i *Inventory) removeProduct(product Product) {
	for ind, p := range i.products {
		if p.id == product.id {
			p.quantity -= product.quantity
			if p.quantity < 0 {
				i.products = append(i.products[:ind], i.products[ind+1:]...)
			}
		}
	}
}

func (i *Inventory) showInventory() {
	fmt.Println("***********************", "Available Products", "***********************", "\n\n")
	for _, p := range i.products {
		fmt.Println(p.name, p.price, p.quantity)
	}
}
