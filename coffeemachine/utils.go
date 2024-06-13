package main

import "fmt"

func printStock(stock *stockSupplies) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", stock.water)
	fmt.Printf("%d ml of milk\n", stock.milk)
	fmt.Printf("%d g of coffee beans\n", stock.beans)
	fmt.Printf("%d disposable cups\n", stock.cups)
	fmt.Printf("$%d of money\n", stock.money)
	fmt.Println()
}

func emptyCashRegister(stock *stockSupplies) string {
	cash := stock.money
	stock.money = 0
	return fmt.Sprintf("Taken $%d from the cash register\n", cash)
}
