package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Prompt the user for input
	fmt.Print("input = ")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Call the Decoder function to process the input
	result := Decoder(input)

	// Display the output
	fmt.Println("output = ", result)
}

func Decoder(input string) string {
	// Convert the input string into an array of symbols
	symbol := make([]string, len(input))
	for i, char := range input {
		symbol[i] = string(char)
	}

	// Create an array to store the numbers
	number := make([]int, len(symbol)+1)

	for i, condition := range symbol {
		if condition == "L" {
			// If 'L' is encountered, adjust the numbers accordingly
			if number[i] <= number[i+1] && i == 0 {
				number[i] = number[i+1] + 1
			} else {
				if number[i] <= number[i+1] {
					number[i] = number[i] + 1
					for j := i; j >= 0; j-- {
						if symbol[j] == "L" {
							number[j] = number[j+1] + 1
						} else if symbol[j] == "=" {
							number[j] = number[j+1]
						}
					}
				}
			}
		} else if condition == "R" {
			// If 'R' is encountered, adjust the numbers accordingly
			if number[i] >= number[i+1] && i == 0 {
				number[i+1] = number[i] + 1
			} else {
				if number[i] >= number[i+1] {
					number[i+1] = number[i] + 1
					for j := i; j >= 0; j-- {
						if symbol[j] == "R" {
							number[j+1] = number[j] + 1
						} else if symbol[j] == "=" {
							number[j] = number[j+1]
						}
					}
				}
			}
		} else if condition == "=" {
			// If '=' is encountered, set the numbers equal
			if i == 0 {
				number[i] = number[i+1]
			} else if number[i] != number[i-1] {
				number[i+1] = number[i]
			}
		}

		fmt.Println(number)
	}

	// Convert the numbers back to a string
	result := ""
	for _, num := range number {
		result += strconv.Itoa(num)
	}

	return result
}
