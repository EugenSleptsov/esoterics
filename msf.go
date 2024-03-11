package esoterics

func MiniStringFuck(input string) string {
	var acc byte
	var output string

	for _, r := range input {
		switch r {
		case '+':
			acc += byte(1)
		case '.':
			output += string(rune(acc))
		}
	}

	return output
}
