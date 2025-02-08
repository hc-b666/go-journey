package main

import (
	"flag"
	"fmt"
	"strings"
)

func formatPhoneNumber(phoneNumber string) (string, error) {
	formatted := strings.ReplaceAll(phoneNumber, " ", "")
	formatted = strings.ReplaceAll(formatted, "-", "")
	formatted = strings.ReplaceAll(formatted, "(", "")
	formatted = strings.ReplaceAll(formatted, ")", "")

	if len(formatted) == 0 {
		return "", nil
	}

	if !strings.HasPrefix(formatted, "+") && len(formatted) != 9 {
		return "", fmt.Errorf("phone number has to start with '+': %s", phoneNumber)
	} else if len(formatted) == 9 {
		digits := "+998"
		for _, ch := range formatted {
			if ch >= '0' && ch <= '9' {
				digits += string(ch)
			}
		}
		formatted = digits
	}

	if !strings.HasPrefix(formatted, "+998") {
		return "", fmt.Errorf("number has to start with +998 only: %s", phoneNumber)
	}
	
	if len(formatted) != 13 {
		return "", fmt.Errorf("total length has to be 13 with '+'")
	}

	return formatted, nil
}

func main() {
	input := flag.String("n", "", "-n=<number>")
	flag.Parse()

	if *input == "" {
		fmt.Println("Enter phone number with -n=<phonenumber>")
		return
	}

	phoneNumbers := strings.Split(*input, ",")

	for _, phoneNumber := range phoneNumbers {
		phoneNumber = strings.TrimSpace(phoneNumber)
		if phoneNumber == "" {
			continue
		}

		formatted, err := formatPhoneNumber(phoneNumber)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}

		if formatted != "" {
			fmt.Println(formatted)
		}
	}
}
