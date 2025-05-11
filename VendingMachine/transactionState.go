package main

type TransactionState int

const (
	BeforeCheckout   = TransactionState(0)
	BeforePayment    = TransactionState(2)
	PaymentInitiated = TransactionState(3)
	ItemDispatched   = TransactionState(4)
)
