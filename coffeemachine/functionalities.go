package main

import "fmt"

func fillStock(stock *stockSupplies) {

	// Prompt user for the initial inputs
	fmt.Println()
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

	fmt.Println()
	fmt.Println("What do you want to buy? 1 - black coffee, 2 - espresso, 3 - latte, 4 - cappuccino, back - to main menu:")
	var option string
	_, _ = fmt.Scan(&option)

	switch option {
	case "1":
		// Black Coffee
		if stock.water >= 100 && stock.beans >= 10 && stock.cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			stock.water -= 100
			stock.beans -= 10
			stock.cups -= 1
			stock.money += 2
		} else if stock.water < 100 {
			fmt.Println("Sorry, not enough water!")
		} else if stock.beans < 10 {
			fmt.Println("Sorry, not enough beans!")
		} else {
			fmt.Println("Sorry, not enough cups!")
		}
	case "2":
		// Espresso
		if stock.water >= 250 && stock.beans >= 16 && stock.cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			stock.water -= 250
			stock.beans -= 16
			stock.cups -= 1
			stock.money += 4
		} else if stock.water < 250 {
			fmt.Println("Sorry, not enough water!")
		} else if stock.beans < 16 {
			fmt.Println("Sorry, not enough beans!")
		} else {
			fmt.Println("Sorry, not enough cups!")
		}
	case "3":
		// Latte
		if stock.water >= 350 && stock.milk >= 75 && stock.beans >= 20 && stock.cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			stock.water -= 350
			stock.milk -= 75
			stock.beans -= 20
			stock.cups -= 1
			stock.money += 7
		} else if stock.water < 350 {
			fmt.Println("Sorry, not enough water!")
		} else if stock.milk < 75 {
			fmt.Println("Sorry, not enough milk!")
		} else if stock.beans < 20 {
			fmt.Println("Sorry, not enough beans!")
		} else {
			fmt.Println("Sorry, not enough cups!")
		}
	case "4":
		// Cappuccino
		if stock.water >= 200 && stock.milk >= 100 && stock.beans >= 12 && stock.cups >= 1 {
			fmt.Println("I have enough resources, making you a coffee!")
			stock.water -= 200
			stock.milk -= 100
			stock.beans -= 12
			stock.cups -= 1
			stock.money += 6
		} else if stock.water < 200 {
			fmt.Println("Sorry, not enough water!")
		} else if stock.milk < 100 {
			fmt.Println("Sorry, not enough milk!")
		} else if stock.beans < 12 {
			fmt.Println("Sorry, not enough beans!")
		} else {
			fmt.Println("Sorry, not enough cups!")
		}
	case "back":
		return
	default:
		fmt.Println("invalid option")
	}
	fmt.Println()
}
