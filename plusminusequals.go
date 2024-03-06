package esoterics

import "strconv"

// PlusMinusEquals is a simple esoteric programming language that uses only three commands: +, -, and =.
// + increments the accumulator, - decrements the accumulator, and = outputs the accumulator as a character and a newline.
func PlusMinusEquals(input string) string {
	var output string
	var accumulator int

	for _, char := range input {
		switch char {
		case '+':
			accumulator++
		case '-':
			accumulator--
		case '=':
			output += strconv.Itoa(accumulator) + "\n"
		}
	}

	return output
}
