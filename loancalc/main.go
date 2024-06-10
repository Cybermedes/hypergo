package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Enter the loan principal:")
	var loanPrincipal int
	_, _ = fmt.Scan(&loanPrincipal)

	menu := `What do you want to calculate?
type "m" for number of monthly payments,
type "p" for the monthly payment:`

	fmt.Println(menu)
	var option string
	_, _ = fmt.Scan(&option)

	switch option {
	case "m":
		calculateNumberOfMonths(loanPrincipal)
	case "p":
		calculateMonthlyPayment(loanPrincipal)
	default:
		fmt.Println("unknown")
	}
}

func calculateNumberOfMonths(principal int) {
	fmt.Println("Enter the monthly payment:")
	var monthlyValue int
	_, _ = fmt.Scan(&monthlyValue)

	monthsToPay := math.Ceil(float64(principal) / float64(monthlyValue))

	// For formatting the final message
	month := "months"
	if monthsToPay == 1 {
		month = "month"
	}

	fmt.Println()
	fmt.Printf("It will take %.0f %s to repay the loan\n", monthsToPay, month)
}

func calculateMonthlyPayment(principal int) {
	fmt.Println("Enter the number of months:")
	var numberOfMonths int
	_, _ = fmt.Scan(&numberOfMonths)

	var monthlyPayment, lastPayment int
	fmt.Println()
	if principal%numberOfMonths == 0 {

		monthlyPayment = principal / numberOfMonths
		fmt.Printf("Your monthly payment = %d\n", monthlyPayment)

	} else {
		value := float64(principal) / float64(numberOfMonths)
		monthlyPayment = int(math.Ceil(value))
		lastPayment = principal - (numberOfMonths-1)*monthlyPayment
		fmt.Printf("Your monthly payment = %d and the last payment = %d\n",
			monthlyPayment, lastPayment)
	}
}
