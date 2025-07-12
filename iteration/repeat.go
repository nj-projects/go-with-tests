package iteration

import "strings"

func Repeat(character string, amount int) string {
	var repeated strings.Builder
	for i := 0; i < amount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()

}
