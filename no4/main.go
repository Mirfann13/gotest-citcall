package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveCryptarithm(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 5 || parts[3] != "=" {
		return "Format input tidak valid"
	}

	letters := make(map[rune]bool)
	for _, part := range parts {
		if part != "+" && part != "-" && part != "=" {
			for _, char := range part {
				if char >= 'A' && char <= 'Z' {
					letters[char] = true
				}
			}
		}
	}

	uniqueLetters := make([]rune, 0, len(letters))
	for letter := range letters {
		uniqueLetters = append(uniqueLetters, letter)
	}

	solution := backtrack(parts, uniqueLetters, make(map[rune]int), 0)
	if solution == nil {
		return "Tidak ada solusi"
	}
	return formatSolution(parts, solution)
}

func backtrack(parts []string, letters []rune, assigned map[rune]int, index int) map[rune]int {
	if index == len(letters) {
		if isValid(parts, assigned) {
			return assigned
		}
		return nil
	}

	for digit := 0; digit <= 9; digit++ {
		if index == 0 && digit == 0 && (strings.HasPrefix(parts[0], string(letters[index])) ||
			strings.HasPrefix(parts[2], string(letters[index])) ||
			strings.HasPrefix(parts[4], string(letters[index]))) {
			continue
		}

		if !isDigitUsed(assigned, digit) {
			newAssigned := make(map[rune]int)
			for k, v := range assigned {
				newAssigned[k] = v
			}
			newAssigned[letters[index]] = digit
			if result := backtrack(parts, letters, newAssigned, index+1); result != nil {
				return result
			}
		}
	}

	return nil
}

func isDigitUsed(assigned map[rune]int, digit int) bool {
	for _, v := range assigned {
		if v == digit {
			return true
		}
	}
	return false
}

func isValid(parts []string, assigned map[rune]int) bool {
	num1 := parseNumber(parts[0], assigned)
	num2 := parseNumber(parts[2], assigned)
	result := parseNumber(parts[4], assigned)

	if num1 == -1 || num2 == -1 || result == -1 {
		return false
	}

	switch parts[1] {
	case "+":
		return num1+num2 == result
	case "-":
		return num1-num2 == result
	default:
		return false
	}
}

func parseNumber(s string, assigned map[rune]int) int {
	result := 0
	for _, char := range s {
		if digit, ok := assigned[char]; ok {
			result = result*10 + digit
		} else {
			return -1
		}
	}
	return result
}

func formatSolution(parts []string, assigned map[rune]int) string {
	var solution strings.Builder
	for i, part := range parts {
		if i > 0 {
			solution.WriteString(" ")
		}
		if part == "+" || part == "-" || part == "=" {
			solution.WriteString(part)
		} else {
			for _, char := range part {
				solution.WriteString(strconv.Itoa(assigned[char]))
			}
		}
	}
	return solution.String()
}

func main() {
	inputs := []string{
		"II + II = HIU",
		"ABD - AD = DKL",
	}

	for _, input := range inputs {
		fmt.Printf("Input: %s\n", input)
		fmt.Printf("Output: %s\n\n", solveCryptarithm(input))
	}
}
