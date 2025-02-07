package main

import (
	"fmt"
)

func main() {
	var deposit, minRate, maxRate float64
	var years int
	var interestType int

	fmt.Print("Deposit amount: ")
	fmt.Scan(&deposit)

	fmt.Print("Min Rate %: ")
	fmt.Scan(&minRate)

	fmt.Print("Max Rate %: ")
	fmt.Scan(&maxRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Interest type (1 - Monthly, 2 - Annually): ")
	fmt.Scan(&interestType)

	minAnnualIncome := deposit * (minRate / 100)
	maxAnnualIncome := deposit * (maxRate / 100)
	minMonthlyIncome := minAnnualIncome / 12
	maxMonthlyIncome := maxAnnualIncome / 12

	fmt.Println("Min Annual Income: ", minAnnualIncome)
	fmt.Println("Max Annual Income: ", maxAnnualIncome)
	fmt.Printf("Min Monthly Income: %.2f \n", minMonthlyIncome)
	fmt.Printf("Max Monthly Income: %.2f \n", maxMonthlyIncome)
}
