package main

import "fmt"

func main() {
	uzsInRuble := 110.00
	uzsInEuro := 13313.00
	uzsInUsd := 12979.39
  
	var amount float64
  var choice int
	var msg string = ""

  fmt.Println("Valuta to UZS Converter")
	
	fmt.Println("1. RUBLE")
	fmt.Println("2. USD")
	fmt.Println("3. EURO")

	fmt.Print("Enter your choice: ")
  fmt.Scan(&choice)

	switch choice {
	
	case 1:
		fmt.Print("Enter amount in RUBLE: ")
		fmt.Scan(&amount)
		result := amount * uzsInRuble
		msg = fmt.Sprintf("%.2f RUBLE = %.2f UZS", amount, result)

	case 2:
		fmt.Print("Enter amount in USD: ")
		fmt.Scan(&amount)
		result := amount * uzsInUsd
		msg = fmt.Sprintf("%.2f USD = %.2f UZS", amount, result)
	
	case 3:
		fmt.Print("Enter amount in EURO: ")
		fmt.Scan(&amount)
		result := amount * uzsInEuro
		msg = fmt.Sprintf("%.2f EURO = %.2f UZS", amount, result)
	
	default:
		fmt.Println("Invalid choice")
	}

  fmt.Println(msg)
}
