package main

import "fmt"

type Transaction struct {
}

func (transaction Transaction) acceptPayment(payment IPayment, amountReqd, amountAdded int) bool {
	if amountAdded < amountReqd {
		fmt.Println("Amount not enough")
		return false
	}
	err := payment.makePayment(amountReqd)
	if err != nil {
		return false
	}
	return true
}

func (transaction Transaction) dispatchItems() {
	// TODO: dispatching items
	fmt.Println("Dispatching Items")
	fmt.Println("Item Dispatched")
}

func (transaction Transaction) dispatchRemainingAmount(amountReqd, amountAdded int) {
	diff := amountAdded - amountReqd
	// TODO: dispatch difference amount
	fmt.Printf("Dispatched %v amount\n", diff)
}

func (transaction Transaction) cancelTransaction(state TransactionState) {
	if state == BeforePayment {
		fmt.Println("Cancelling payment, redirect to Cart page")
	} else if state == PaymentInitiated {
		fmt.Println("Cancelling payment, any amount deucted will be refunded")
	} else if state == ItemDispatched {
		fmt.Println("Can't cancel order, item will be dispatched")
	}
}
