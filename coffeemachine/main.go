package main

import (
	"fmt"
)

type stockSupplies struct {
	water int
	milk  int
	beans int
	cups  int
	money int
}

func main() {

	// Initial stats of our program
	coffeeMachine := stockSupplies{
		water: 400,
		milk:  540,
		beans: 120,
		cups:  9,
		money: 550,
	}

	printStock(&coffeeMachine)
	fmt.Println("Write action (buy, fill, take):")
	var action string
	_, _ = fmt.Scan(&action)

	switch action {
	case "buy":
		buyCoffee(&coffeeMachine)
	case "fill":
		fillStock(&coffeeMachine)
	case "take":
		fmt.Println(emptyCashRegister(&coffeeMachine))
	default:
		fmt.Println("invalid option, write buy, fill or take")
	}
	printStock(&coffeeMachine)
}

func printStock(stock *stockSupplies) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", stock.water)
	fmt.Printf("%d ml of milk\n", stock.milk)
	fmt.Printf("%d g of coffee beans\n", stock.beans)
	fmt.Printf("%d disposable cups\n", stock.cups)
	fmt.Printf("$%d of money\n", stock.money)
	fmt.Println()
}

func fillStock(stock *stockSupplies) {
	// Prompt user for the initial inputs
	fmt.Println("Write how many ml of water you want to add:")
	var water int
	_, _ = fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	var milk int
	_, _ = fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	var beans int
	_, _ = fmt.Scan(&beans)
	fmt.Println("Write how many disposable cups you want to add:")
	var numberOfCups int
	_, _ = fmt.Scan(&numberOfCups)

	stock.water += water
	stock.milk += milk
	stock.beans += beans
	stock.cups += numberOfCups
	fmt.Println()
}

func buyCoffee(stock *stockSupplies) {

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var option int
	_, _ = fmt.Scan(&option)

	switch option {
	case 1:
		// Espresso
		if stock.water >= 250 && stock.beans >= 16 && stock.cups >= 1 {
			stock.water -= 250
			stock.beans -= 16
			stock.cups -= 1
			stock.money += 4
		}
	case 2:
		// Latte
		if stock.water >= 350 && stock.milk >= 75 && stock.beans >= 20 && stock.cups >= 1 {
			stock.water -= 350
			stock.milk -= 75
			stock.beans -= 20
			stock.cups -= 1
			stock.money += 7
		}
	case 3:
		// Cappuccino
		if stock.water >= 200 && stock.milk >= 100 && stock.beans >= 12 && stock.cups >= 1 {
			stock.water -= 200
			stock.milk -= 100
			stock.beans -= 12
			stock.cups -= 1
			stock.money += 6
		}
	default:
		fmt.Println("invalid option")
	}
	fmt.Println()
}

func emptyCashRegister(stock *stockSupplies) string {
	cash := stock.money
	stock.money = 0
	return fmt.Sprintf("I gave you $%d", cash)
}
