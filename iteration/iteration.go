package iteration

func Repeat(character string, amount int) string {
	var s string
	for i := 0; i < amount; i++ {
		s += character
	}
	return s
}
