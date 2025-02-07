package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("n", "User", "-n <name>")
	code := flag.Int("c", 0, "-c <code>")
	lang := flag.String("l", "en", "-l <lang>")

	flag.Parse()

	if *name == "User" {
		fmt.Println("You have to specify the name of the user")
		return
	}

	if *code == 0 {
		fmt.Println("You have to pass the code")
		return
	}

	if *lang == "en" {
		fmt.Printf("Welcome, %s! Your access code is: %d", *name, *code)
	} else if *lang == "ru" {
		fmt.Printf("Добро пожаловать, %s! Ваш код доступа: %d", *name, *code)
	}
}
