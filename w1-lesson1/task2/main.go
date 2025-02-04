package main

import (
	"fmt"
	"strings"
)

const (
	uzsInRuble = 110.00
	uzsInEuro  = 13313.00
	uzsInUsd   = 12979.39
)

type Currency struct {
	Name  string
	Rate  float64
	Label string
}

var currencies = map[string]Currency{
	"1": {
		"RUBLE",
		uzsInRuble,
		"RUBLE",
	},
	"2": {
		"USD",
		uzsInUsd,
		"USD",
	},
	"3": {
		"EURO",
		uzsInEuro,
		"EURO",
	},
}

func logMenu() {
	fmt.Println("====================")
	fmt.Println("Valuta to UZS Converter")

	for key, currency := range currencies {
		fmt.Printf("%s. %s\n", key, currency.Name)
	}

	fmt.Println("x. Exit")
	fmt.Println("====================")
}

func getAmount(currencyLabel string) float64 {
	var amount float64
	fmt.Printf("Enter amount in %s: ", currencyLabel)
	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println("Invalid input. Please enter a numeric value.")
		return 0
	}
	return amount
}

func main() {
	isRunning := true

	for isRunning {
		logMenu()

		var choice string
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		choice = strings.ToLower(choice)

		if choice == "x" {
			isRunning = false
			fmt.Println("Exiting...")
			continue
		}

		currency, exists := currencies[choice]
		if !exists {
			fmt.Println("Invalid choice")
			continue
		}

		amount := getAmount(currency.Label)
		if amount <= 0 {
			continue
		}

		result := amount * currency.Rate
		msg := fmt.Sprintf("%.2f %s = %.2f UZS", amount, currency.Label, result)
		fmt.Println(msg)
	}
}
