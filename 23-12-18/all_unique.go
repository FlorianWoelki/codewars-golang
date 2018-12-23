package main

import (
	"fmt"
	"strings"
)

func HasUniqueChar(str string) bool {
	for _, char := range str {
		if strings.Count(str, string(char)) != 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(HasUniqueChar("  nAa")) // false
	fmt.Println(HasUniqueChar("abcde")) // true
	fmt.Println(HasUniqueChar("++-")) // false
	fmt.Println(HasUniqueChar("AaBbC")) // true
}
