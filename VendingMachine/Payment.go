package main

import "errors"

type IPayment interface {
	makePayment(amount int) error
}

type UPIPayment struct {
}

func (upip UPIPayment) makePayment(amount int) error {
	//TODO: Make Payment via UPI
	return nil
}

type CreditCart struct {
}

func (cc CreditCart) makePayment(amount int) error {
	// TODO: Make Payment via Card
	return errors.New("random error")
}
