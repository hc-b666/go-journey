package main

import "fmt"

func main() {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	msg := fmt.Sprintf("Welcome to Alif Academy, %s! Learning Go is exciting!", name)

	fmt.Println(msg)
}
