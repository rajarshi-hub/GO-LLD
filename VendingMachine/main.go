package main

func main() {
	inventory := Inventory{
		products: []Product{},
	}
	cart := Cart{
		products: []Product{},
	}
	trans := Transaction{}
	transState := TransactionState(0)

	machine := VendingMachine{
		inventory:        inventory,
		Cart:             cart,
		transaction:      trans,
		transactionState: transState,
	}
	prod1 := Product{
		name:     "Grocery",
		price:    10,
		id:       "123",
		quantity: 12,
	}
	prod2 := Product{
		name:     "Milk",
		price:    20,
		id:       "456",
		quantity: 20,
	}
	prod3 := Product{
		name:     "Uncle Chips",
		price:    30,
		id:       "789",
		quantity: 30,
	}
	machine.addInventory(prod1)
	machine.addInventory(prod2)
	machine.addInventory(prod3)

	prod1 = Product{
		name:     "Grocery",
		price:    10,
		id:       "123",
		quantity: 2,
	}

	prod2 = Product{
		name:     "Milk",
		price:    20,
		id:       "456",
		quantity: 3,
	}
	machine.addProductToCart(prod1)
	machine.addProductToCart(prod2)
	machine.checkout()
	// Normal flow
	//upi := UPIPayment{}
	//machine.makePayment(upi, 90)

	// Cancel Transaction flow
	//machine.cancelTransaction()

	// transaction unsuccessful
	credit := CreditCart{}
	machine.makePayment(credit, 90)

}
