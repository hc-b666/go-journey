package main

import "fmt"

func main() {
	var totalPrice float64
	var bonus float64
	var discount float64

	fmt.Print("Enter Total Price: ")
	fmt.Scan(&totalPrice)

	fmt.Print("Enter Bonus in %: ")
	fmt.Scan(&bonus)

	fmt.Print("Enter Discount in %: ")
	fmt.Scan(&discount)

	result := (bonus / 100) * (totalPrice * (100 - discount) / 100)

	msg := fmt.Sprintf("Bonus is %.2f", result)
	fmt.Println(msg)
}
