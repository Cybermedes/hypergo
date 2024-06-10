package main

import (
	"fmt"
)

func main() {

	// Prompt user for the initial inputs
	fmt.Println("Write how many ml of water the coffee machine has:")
	var water int
	_, _ = fmt.Scan(&water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	var milk int
	_, _ = fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	var beans int
	_, _ = fmt.Scan(&beans)
	fmt.Println("Write how many cups of coffee you will need:")
	var numberOfCups int
	_, _ = fmt.Scan(&numberOfCups)

	fmt.Println(checkIngredients(water, milk, beans, numberOfCups))
}

func checkIngredients(water, milk, beans, quantity int) string {

	var cups int
	for {
		if water >= 200 && milk >= 50 && beans >= 15 {
			water -= 200
			milk -= 50
			beans -= 15
			cups++
		} else {
			break
		}
	}

	if quantity == cups {
		return "Yes, I can make that amount of coffee"
	} else if quantity < cups {
		return fmt.Sprintf("Yes, I can make that amount of coffee "+
			"(and even %d more than that)", cups-quantity)
	} else {
		return fmt.Sprintf("No, I can make only %d cups of coffee", cups)
	}
}
