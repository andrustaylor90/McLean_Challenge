package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// validateCreditCard checks if the given credit card number is valid according to specified rules.
func validateCreditCard(number string) bool {
	// Regex to check the structure: starts with 4, 5, or 6; digits can be in groups of four separated by a single hyphen.
	regex := regexp.MustCompile(`^([4-6]\d{3}-\d{4}-\d{4}-\d{4}|[4-6]\d{15})$`)

	// Check for overall pattern match
	if !regex.MatchString(number) {
		return false
	}

	// Remove hyphens for further validation
	cleanNumber := strings.Replace(number, "-", "", -1)

	// Check for four or more consecutive repeated digits manually
	return !hasFourConsecutiveSameDigits(cleanNumber)
}

// hasFourConsecutiveSameDigits checks if there are four or more consecutive identical digits in a string.
func hasFourConsecutiveSameDigits(s string) bool {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count == 4 {
				return true
			}
		} else {
			count = 1
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line for the number of credit card numbers to be validated
	scanner.Scan()
	n := 0
	fmt.Sscanf(scanner.Text(), "%d", &n)

	// Iterate over each credit card number and validate
	results := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cardNumber := scanner.Text()
		if validateCreditCard(cardNumber) {
			results[i] = "Valid"
		} else {
			results[i] = "Invalid"
		}
	}

	// Output the results
	for _, result := range results {
		fmt.Println(result)
	}
}
