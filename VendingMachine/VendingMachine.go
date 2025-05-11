package main

import "fmt"

type VendingMachine struct {
	inventory Inventory // We can promote these to top level as well
	Cart
	transaction      Transaction
	transactionState TransactionState
}

func (vm *VendingMachine) showInventory() {
	vm.inventory.showInventory()
}

func (vm *VendingMachine) showCart() {
	vm.Cart.showCart()
}

func (vm *VendingMachine) addProductToCart(pr Product) {
	vm.AddProduct(pr)
	vm.showCart()
	vm.showInventory()
}

func (vm *VendingMachine) removeProductFromCart(pr Product) {
	vm.Cart.removeProduct(pr)
}

func (vm *VendingMachine) addInventory(pr Product) {
	vm.inventory.addProduct(pr)
	vm.showInventory()
}

func (vm *VendingMachine) checkout() {
	fmt.Println("Checking out for transaction, redirect to transaction state")
	vm.transactionState = BeforePayment
}

func (vm *VendingMachine) makePayment(payment IPayment, amount int) {
	vm.transactionState = PaymentInitiated
	amount2 := vm.Cart.getCartTotal()
	done := vm.transaction.acceptPayment(payment, amount2, amount)

	if done {
		vm.transaction.dispatchItems()
		// TODO: Remove items from Inventoru
		vm.transactionState = ItemDispatched
		vm.transaction.dispatchRemainingAmount(amount2, amount)
	} else {
		vm.transaction.cancelTransaction(vm.transactionState)
	}
}

func (vm *VendingMachine) dispatchItems() {

}

func (vm *VendingMachine) cancelTransaction() {
	vm.transaction.cancelTransaction(vm.transactionState)
}
