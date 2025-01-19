package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Function to read a multi-line template from user input
func readTemplate() string {
	fmt.Println("Enter your prompt (type END on a new line to finish):")

	scanner := bufio.NewScanner(os.Stdin)
	var builder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "END" {
			break
		}
		builder.WriteString(line + "\n")
	}

	// checks for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input", err)
		os.Exit(1)
	}

	return builder.String()
}

// function to extract placeholders
func extractPlaceholders(s string) [][]string {
	re := regexp.MustCompile(`\[(.*?)\]`)
	return re.FindAllStringSubmatch(s, -1)

}

// function to get user inputs for the placeholders
func getUserInput(placeholders [][]string) []string {
	var inputs []string
	scanner := bufio.NewScanner(os.Stdin)

	for _, placeholder := range placeholders {
		input := ""

		for input == "" { // keep asking until valid input
			fmt.Println(placeholder[1], ":")

			scanner.Scan()
			input = strings.TrimSpace(scanner.Text())
			if input == "" {
				fmt.Println("Input cannot be empty. Please try again.")
			}
		}

		inputs = append(inputs, input)
	}

	// checks for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input", err)
		os.Exit(1)
	}

	return inputs
}

// Function to replace placeholders with user inputs
func replacePlaceholders(s string, placeholders [][]string, inputs []string) string {
	for idx, val := range inputs {
		s = strings.Replace(s, placeholders[idx][0], val, 1)
	}
	return s
}

func displayResult(s string) {
	fmt.Println("Filled prompt:\n\n", s)
	fmt.Println("\n\n --------Press enter to exit--------")
	fmt.Scanln()
}

func main() {

	tempStr := readTemplate()
	matches := extractPlaceholders(tempStr)
	inputs := getUserInput(matches)

	tempStr = replacePlaceholders(tempStr, matches, inputs)

	displayResult(tempStr)

}
