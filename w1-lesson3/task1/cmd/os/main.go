package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: ./task1 <name> <code>")
		return
	}

	name := os.Args[1]
	code, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("<code> must be interger")
		return;
	}
	lang := os.Args[3]

	if lang == "en" {
		fmt.Printf("Welcome, %s! Your access code is: %d", name, code)
	} else if lang == "ru" {
		fmt.Printf("Добро пожаловать, %s! Ваш код доступа: %d", name, code)
	} else {
		fmt.Printf("Welcome, %s! Your access code is: %d", name, code)
	}
}
