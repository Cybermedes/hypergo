package main

import (
	"fmt"
	"os"
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
		water: 2000,
		milk:  1000,
		beans: 250,
		cups:  12,
		money: 150,
	}

	for {

		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		_, _ = fmt.Scan(&action)
		fmt.Println()

		switch action {
		case "buy":
			buyCoffee(&coffeeMachine)
		case "fill":
			fillStock(&coffeeMachine)
		case "take":
			fmt.Println(emptyCashRegister(&coffeeMachine))
		case "remaining":
			printStock(&coffeeMachine)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("invalid option, write buy, fill, take, remaining or exit")
		}
	}
}
