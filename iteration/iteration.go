package iteration

const count = 5

// Repeat the string count times
func Repeat(character string) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return
}
