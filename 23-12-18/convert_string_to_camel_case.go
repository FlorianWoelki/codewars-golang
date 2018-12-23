package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	splittedString := strings.Split(s, "-")

	if len(splittedString) == 1 {
		splittedString = strings.Split(s, "_")
	}

	resultString := splittedString[0] + " "
	for i := 1; i < len(splittedString); i++ {
		resultString += strings.Title(splittedString[i]) + " "
	}

	return strings.Replace(resultString, " ", "", -1)
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}